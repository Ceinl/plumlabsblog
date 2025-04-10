package articles

import (
	"database/sql"
	"io"
	"mime/multipart"
	"time"
)

type Manager struct {
	db        *sql.DB
	articles  []Article
}

type Article struct {
	Title       string
	mdContent   string
	htmlContent string
	Edited		bool
	Created     string
	LastUpdate  string
}

func NewArticleManager(db *sql.DB) *Manager{
	return &Manager{
		db: db,
	}
}

func (m Manager) Handle(file *multipart.FileHeader) error {
	name, extention := splitName(file)
	if !m.isArticleExist(name) && extention == ".md"{
		article, err := m.CreateArticle(file)	
		if err != nil {
			return err
		}		
		m.articles = append(m.articles, article)
	}
	return nil
}

func (m Manager) CreateArticle(file *multipart.FileHeader) (Article,error){
	var article Article

	article.Title, _ = splitName(file)

	err := article.GetContent(file)
	if err != nil { return article,err }

	err = article.ConvertToHTML()
	if err != nil { return article,err }

	article.Edited = false

	article.Created = time.Now().Format(time.RFC3339)
	article.LastUpdate = time.Now().Format(time.RFC3339)

	return article,nil
}

func (m Manager) isArticleExist(title string) bool { 
	for _,article := range m.articles{
		if article.Title == title{
			return true
		}
	}
	return false
}

func splitName(file *multipart.FileHeader) (string,string) {
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

func (a Article) GetContent(fileheader *multipart.FileHeader) error {
	file, err := fileheader.Open()
	if err != nil {
		return err
	}
	defer file.Close() 

	content , err := io.ReadAll(file) 
	if err != nil {
		return err
	}

	a.mdContent = string(content)
	return nil 
}

func (a Article) ConvertToHTML() error {
	return nil
}


