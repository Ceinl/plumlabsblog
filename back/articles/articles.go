package articles

import (
	"database/sql"
	"io"
	"mime/multipart"
	"plumlabs/back/utils/manager"
)

type Manager struct {
	db        *sql.DB
	Articles  []Article
}

type Article struct {
	Id          int
	Title       string
	MdContent   string
	HtmlContent string
	Edited		bool
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
		if err != nil { return err }		
		m.Articles = append(m.Articles, article)
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

	return article,nil
}

func (m Manager) isArticleExist(title string) bool { 
	for _,article := range m.Articles{
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

	a.MdContent = string(content)
	return nil 
}

func (a Article) ConvertToHTML() error {
	content, err := manager.ArticleManage(a.MdContent)
	if err != nil { return err}
	a.HtmlContent = content
	return nil
}


