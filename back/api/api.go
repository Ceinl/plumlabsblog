package api

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	article_manager "plumlabs/back/articles"
)

type ArticleData struct {
	HTMLContent template.HTML 
}

type ArticlesData struct {
	Titles []string
}

type API struct {
	db             *sql.DB
	articleManager article_manager.Manager
	templates      *template.Template 
}

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func New(db *sql.DB) API {
	api := API{db: db}

	am := *article_manager.NewArticleManager(db)
	api.articleManager = am

	tmpl, err := template.New("").Funcs(template.FuncMap{"safeHTML": safeHTML}).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err) 
	}
	api.templates = tmpl
	log.Println("HTML templates loaded successfully.")

	return api
}

func (api *API) ApiPostFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "HX-Request, HX-Target, HX-Current-URL, Content-Type")

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

	err = api.articleManager.Handle(handler)
	if err != nil {
		http.Error(w, "<div class='error'>Failed to process article: "+err.Error()+"</div>", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	log.Printf("Article uploaded successfully")
	w.Write([]byte("<div id='upload-result' class='success'>Article uploaded successfully</div>"))
}

func (api *API) ApiDeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "HX-Request, HX-Target, HX-Current-URL, Content-Type")

	if r.Method == http.MethodOptions {
		log.Printf("OPTIONS request received")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		log.Printf("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "HX-Request, HX-Target, HX-Current-URL, Content-Type")

	title := r.URL.Query().Get("title")

	if title == "" {
		http.Error(w, "Missing title parameter", http.StatusBadRequest)
		return
	}
	htmlContent, err := api.articleManager.ReadHtmlArticle(title)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")

	data := ArticleData{HTMLContent: template.HTML(htmlContent)} 
	err = api.templates.ExecuteTemplate(w, "article.html", data)
	if err != nil {
		log.Printf("Error executing article template: %v", err)
		http.Error(w, "Failed to render article", http.StatusInternalServerError)
	}
}

func (api *API) ApiGetTitles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "HX-Request, HX-Target, HX-Current-URL, Content-Type")

	titles, err := api.articleManager.ReadAllArticleTitles()
	if err != nil {
		log.Printf("Error getting titles: %v", err)
		http.Error(w, "<div class='error'>Failed to get titles: "+err.Error()+"</div>", http.StatusInternalServerError)
		return
	}

	if titles == nil {
		titles = []string{}
	}

	w.Header().Set("Content-Type", "text/html")

	data := ArticlesData{Titles: titles}
	err = api.templates.ExecuteTemplate(w, "articles.html", data)
	if err != nil {
		log.Printf("Error executing articles template: %v", err)
		http.Error(w, "Failed to render articles list", http.StatusInternalServerError)
	}
}

