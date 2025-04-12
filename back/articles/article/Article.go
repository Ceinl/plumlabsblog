package Article

import (
	"io"
	"log"
	"mime/multipart"
	"plumlabs/back/utils/manager"
)

type Article struct {
	Id          int
	Title       string
	MdContent   string
	HtmlContent string
}

func (a *Article) GetContent(fileheader *multipart.FileHeader) error {
	log.Printf("Getting content from file: %s", fileheader.Filename)
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

func (a *Article) ConvertToHTML() error {
	log.Printf("Converting to HTML")
	content, err := manager.ArticleManage(a.MdContent)
	if err != nil { return err}
	a.HtmlContent = content
	return nil
}
