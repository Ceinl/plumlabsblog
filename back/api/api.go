package api

import (
	"database/sql"
	"log"
	"net/http"
	article_manager "plumlabs/back/articles"
)

type API struct {
	db             *sql.DB
	articleManager article_manager.Manager
}

func New(db *sql.DB) API {
	api := API{db: db}

	am := *article_manager.NewArticleManager(db)
	api.articleManager = am

	return api
}

func (api *API) ApiPostFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "<div class='error'>Method not allowed</div>", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) 
	if err != nil {
		http.Error(w, "<div class='error'>Failed to parse form: "+err.Error()+"</div>", http.StatusBadRequest)
		return
	}

	_, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "<div class='error'>Error retrieving file: "+err.Error()+"</div>", http.StatusBadRequest)
		return
	}

	log.Printf(handler.Filename)
	err = api.articleManager.Handle(handler)
	if err != nil {
		http.Error(w, "<div class='error'>Failed to process article: "+err.Error()+"</div>", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<div id='upload-result' class='success'>Article uploaded successfully</div>"))
}

func (api *API) ApiDeleteArticle(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Query().Get("title")
    if title == "" {
        http.Error(w, "<div class='error'>Missing title parameter</div>", http.StatusBadRequest)
        return
    }

    api.articleManager.DeleteArtile(title)

    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("<div id='delete-result' class='success'>Article '" + title + "' deleted successfully</div>"))
}

func (api *API) ApiGetArticle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	html, err := api.articleManager.ReadHtmlArticle(title)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func (api *API) ApiGetTitles(w http.ResponseWriter, r *http.Request) {
	titles, err := api.articleManager.ReadAllArticleTitles()
	if err != nil {
		log.Printf("Error getting titles: %v", err)
		http.Error(w, "<div class='error'>Failed to get titles: "+err.Error()+"</div>", http.StatusInternalServerError)
		return
	}
	
	// Check if titles is nil or empty
	if titles == nil {
		titles = []string{} // Initialize with empty slice to avoid nil pointer
	}
	
	w.Header().Set("Content-Type", "text/html")
	html := "<ul>"
	for _, title := range titles {
		html += "<li>" + title + "</li>"
	}
	html += "</ul>"
	w.Write([]byte(html))
}
