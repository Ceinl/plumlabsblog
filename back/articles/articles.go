package articles

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"plumlabs/back/utils/manager"
	"strings"
)

// HandleFile processes a Markdown file, validates its extension, and creates
// both .md and .html versions in the storage directory.
// It returns an error if the file is not a valid Markdown file or if any
// processing step fails.
func HandleFile(files *multipart.FileHeader) error {

	name, extention, err:= splitName(files)
	if err != nil {
		return err
	}

	if extention != "md"{
		return errors.New("incorrect file extention")
	}

	if !isDirExist(name) {
		err = CreateArticle(name, files)

		if err != nil{
			return err
		}
	}

	return nil
}

// isDirExist checks if a directory with the given name exists in the articles storage path.
// It returns true if the directory exists, false otherwise.
func isDirExist(name string) bool{
	dirpath := filepath.Join("storage", "articles", name)
	info, err := os.Stat(dirpath)

	if os.IsNotExist(err){
		return false 
	}
	return err == nil && info.IsDir() 
}

// splitName extracts the name and extension from a file.
// It returns the name, extension, and an error if the filename format is invalid.
func splitName(file *multipart.FileHeader) (string, string, error) {
	filename :=	file.Filename
	splited := strings.Split(filename, ".")
	if len(splited) != 2{
		return "", "", errors.New("incorrect file name")	
	}
	return splited[0] , splited[1], nil
}

// CreateArticle creates a new article directory with content.md and content.html files.
// It returns an error if any step in the creation process fails.
func CreateArticle(name string, fh *multipart.FileHeader) error {
	err := createMd(name,fh)
	if err != nil{ return err }

	err = createHTML(name)
	if err != nil{ return err }

	return nil

}
// ReadArticle retrieves the content of an article file from storage.
// The extension parameter should include the dot, e.g., ".md" or ".html".
// It returns the file content as a string and any error encountered.
func ReadArticle(name, extention string) (string, error) {
	dirpath := filepath.Join("storage", "articles", name)
	filepath := filepath.Join(dirpath,name + extention)

	data, err := os.ReadFile(filepath)
	if err != nil { return "",err }
	return string(data),nil	
}
// UpdateArticle updates an existing article's content.
// Not implemented yet.
func UpdateArticle() {
	// TODO: Implement article update functionality
}

// DeleteArticle removes an article from storage.
// Not implemented yet.
func DeleteArticle() {
	// TODO: Implement article deletion functionality
}

// createMd creates a content.md file in the article's directory from the uploaded file.
// It returns an error if any step in the creation process fails.
func createMd(name string, fh *multipart.FileHeader) error {
	dirpath := filepath.Join("storage", "articles", name)
	
	err := os.MkdirAll(dirpath, 0755)
	if err != nil { return err }

	file, err := fh.Open()	
	if err != nil{ return err }
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil{ return err }

	out, err := os.Create(filepath.Join(dirpath, "content" + ".md"))
	if err != nil { return err }
	defer out.Close()

	out.Write(data)

	return nil
}

// createHTML reads the content.md file, converts it to HTML using the markdown manager,
// and saves the result as content.html in the same directory.
// It returns an error if any step in the process fails.
func createHTML(name string) error { 

	dirpath := filepath.Join("storage", "articles", name)
	filepath := filepath.Join(dirpath,"content.md")

	data, err := os.ReadFile(filepath)
	if err != nil { return err }

	html, err := manager.ArticleManage(string(data)) 
	if err != nil { return err }

	out, err := os.Create(dirpath + "/content.html")
	if err != nil{ return err }

	out.Write([]byte(html))


	return nil
}
