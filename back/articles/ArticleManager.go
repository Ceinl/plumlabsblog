package article_manager

import (
	"database/sql"
	"log"
	"mime/multipart"
	Article "plumlabs/back/articles/article"
	"plumlabs/back/storage"
	"slices"
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
		if extention == "md" {
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

/*
// CRUD 
//
// CreateArticle takes file return Article and error
//
// OTHER OPERATIONS
//
//
*/
func (m *Manager) CreateArticle(file *multipart.FileHeader) (Article.Article,error){
	log.Printf("Creating article from file: %s", file.Filename)
	var article Article.Article

	article.Title, _ = splitName(file)

	err := article.GetContent(file)
	if err != nil { return article,err }

	err = article.ConvertToHTML()
	if err != nil { return article,err }

	return article,nil
}

func (m *Manager) ReadHtmlArticle(title string) (content string, err error) {
	err = nil

	a, err := storage.GetArticleByTitle(m.db,title)
	if err != nil { return }
	content = a.HtmlContent

	return 
}


func (m *Manager) ReadMdArticle() (content string,title string, err error) {
	err = nil

	a, err := storage.GetArticleByTitle(m.db,title)
	if err != nil { return }
	content = a.MdContent

	return 
}

func (m *Manager) ReadAllArticleTitles() (titles []string, err error) {
	err = nil

	articles, err := storage.GetAllArticles(m.db)

	for _ , article := range articles{
		titles = append(titles, article.Title)
	}

	if err != nil { return }

	return
}

func (m *Manager) UpdateArticle(title string, file *multipart.FileHeader) error { 
	log.Printf("Updating article")
	article, err:= m.CreateArticle(file)
	if err != nil {return err}
	storage.UpdateAricle(m.db,article) 
	return nil
}

func (m *Manager) DeleteArtile(title string) {
	log.Printf("Deleting article")
	storage.DeleteArticle(m.db,title)
	for i,article := range m.Articles{
		if article.Title == title{
			m.Articles = slices.Delete(m.Articles, i, i+1)
		}
	}
}


func (m *Manager) isArticleExist(title string) bool { 
	log.Printf("Checking if article exist: %s", title)
	_, err := storage.GetArticleByTitle(m.db,title)
	return err == nil
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



