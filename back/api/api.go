package api

import (
	"database/sql"
	//	"mime/multipart"
	"net/http"
	article_manager "plumlabs/back/articles"
)

type API struct {
	db        *sql.DB
	articleManager article_manager.Manager	
}

func New(db *sql.DB) API {
	api := API{ db: db }

	am := *article_manager.NewArticleManager(db)
	api.articleManager= am

	return api
}

func (api *API) ApiPostFile(w http.ResponseWriter, r *http.Request) {
	
	// Extract file from request and send it to handle

	api.articleManager.Handle(nil)
}

func (api *API) ApiDeleteArticle(w http.ResponseWriter, r *http.Request) { 

	// Extract string from request and sent it to method

	api.articleManager.DeleteArtile("")

}

func (api *API) ApiGetArticle(w http.ResponseWriter, r *http.Request) { 

	// Extract string from request and sent it to read method. Return content wrap with a template and send to front

	api.articleManager.ReadHtmlArticle("")

}

func (api *API) ApiGetTitles(w http.ResponseWriter, r *http.Request) { 

	//Return content wrap with a template and send to front

	api.articleManager.ReadAllArticleTitles()

}
