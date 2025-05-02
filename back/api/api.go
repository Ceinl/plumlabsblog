package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	html, err := api.articleManager.ReadHtmlArticle(title)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")

	// TODO: return not html content but tamplate with html as param
	w.Write(ArticleWrapper(html))
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

	// Use the updated AllArticlesWrapper
	w.Write(AllArticlesWrapper(titles))
}

// Updated AllArticlesWrapper to use Tailwind and HTMX
func AllArticlesWrapper(titles []string) []byte {
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<div id='all-articles' class='p-4 bg-gray-800 rounded-lg shadow-md'>") // Container with Tailwind styling
	htmlBuilder.WriteString("<h3 class='text-xl font-semibold mb-3 text-white'>Available Articles</h3>")
	if len(titles) == 0 {
		htmlBuilder.WriteString("<p class='text-gray-400'>No articles found.</p>")
	} else {
		htmlBuilder.WriteString("<ul class='space-y-2'>")
		for _, title := range titles {
			// Added text-center and changed text color classes
			htmlBuilder.WriteString(fmt.Sprintf(
				`<li
                    class='text-green-400 hover:text-green-300 cursor-pointer p-2 rounded hover:bg-gray-700 transition duration-150 ease-in-out text-center'
					hx-get="http://localhost:1612/api/article/get?title=%s"
                    hx-target="#article-display"
                    hx-swap="innerHTML"
                    hx-indicator="#loading-indicator">
                    %s
                 </li>`,
				url.QueryEscape(title), // Ensure title is URL-encoded
				title,
			))
		}
		htmlBuilder.WriteString("</ul>")
	}
	// Added a loading indicator for HTMX requests
	htmlBuilder.WriteString("<div id='loading-indicator' class='htmx-indicator text-gray-400 mt-2'>Loading...</div>")
	htmlBuilder.WriteString("</div>") // Close container div
	// Added a target div for displaying the selected article
	htmlBuilder.WriteString("<div id='article-display' class='mt-4'></div>")

	return []byte(htmlBuilder.String())
}

// Updated ArticleWrapper to use Tailwind
func ArticleWrapper(html string) []byte {
	// Added Tailwind classes for styling the article content area
	wrapperHtml := fmt.Sprintf("<div id='article' class='prose prose-invert max-w-none p-4 bg-gray-700 rounded-lg shadow-inner'>%s</div>", html)
	return []byte(wrapperHtml)
}
