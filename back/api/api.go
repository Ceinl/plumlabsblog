package api

import (
	"database/sql"
//	"mime/multipart"
	"net/http"
	"plumlabs/back/articles"
)

type API struct {
	db        *sql.DB
	am		  article_manager.Manager	
}

func New(db *sql.DB) *API {
	api := API{ db: db }

	am := *article_manager.NewArticleManager(db)
	api.am = am

	return &api
}

func (api *API) ApiPostFile(w http.ResponseWriter, r *http.Request) {}

func (api *API) ApiDeleteArticle(w http.ResponseWriter, r *http.Request) { }

func (api *API) ApiGetArticle(w http.ResponseWriter, r *http.Request) { }

func (api *API) ApiGetTitles(w http.ResponseWriter, r *http.Request) { }
