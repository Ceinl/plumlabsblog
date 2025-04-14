package api

import (
	"database/sql"
	"mime/multipart"
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

func (api *API) ApiPostFile(file *multipart.FileHeader) error {
	return nil
}

func (api *API) ApiDeleteArticle(title string) error {
	return nil
}

func (api *API) ApiGetArticle(title string) string {
	return ""
}

func (api *API) ApiGetTitles() string{
	return ""

}
