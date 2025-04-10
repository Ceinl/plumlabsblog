package article_manager

import (
	"database/sql"
	"log"
	"mime/multipart"
	Article "plumlabs/back/articles/article"
	"plumlabs/back/storage"
)

type Manager struct {
	db        *sql.DB
	Articles  []Article.Article
}


func NewArticleManager(db *sql.DB) *Manager{
	log.Printf("ArticleManager created")
	return &Manager{
		db: db,
	}
}

func (m *Manager) Handle(file *multipart.FileHeader) error {
	log.Printf("file received: %s",file.Filename)
	name, extention := splitName(file)
	if !m.isArticleExist(name) {
		log.Printf("Creating new article")
		if extention == ".md" {
			article, err := m.CreateArticle(file)	
			if err != nil { return err }		
			m.Articles = append(m.Articles, article)
		}else{
			log.Printf("Wrong extention")
		}
	}else {
		log.Printf("Updating old article")
	
	}
	return nil
}

func (m *Manager) CreateArticle(file *multipart.FileHeader) (Article.Article,error){
	log.Printf("Creating article from file: %s", file.Filename)
	var article Article.Article

	article.Title, _ = splitName(file)

	err := article.GetContent(file)
	if err != nil { return article,err }

	err = article.ConvertToHTML()
	if err != nil { return article,err }

	article.Edited = false

	return article,nil
}

func (m *Manager) UpdateArticle(title string, file *multipart.FileHeader) (*Article.Article,error){
	log.Printf("Updating article")
	article , err := storage.GetArticleByTitle(m.db, title) 
	if err != nil {return article, err}	

	return article,nil
}


func (m *Manager) isArticleExist(title string) bool { 
	log.Printf("Checking if article exist: %s", title)
	for _,article := range m.Articles{
		if article.Title == title{
			return true
		}
	}
	return false
}

func splitName(file *multipart.FileHeader) (string,string) {
	log.Printf("Splitting filename: %s", file.Filename)
	name, extention := "", ""
	for i := len(file.Filename)-1; i >=0 ; i-- {
		if string(file.Filename[i]) == "."{
			name = file.Filename[:i]
			extention = file.Filename[i+1:]
			return name, extention
		}
	}
	return name, extention
}



