package articles

import (
	"archive/zip"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"plumlabs/back/utils/manager"
	"strings"
)

// ArticleManager handles article operations including file processing and storage
type ArticleManager struct {
	basePath string
	article  Article
}

// Article represents a blog article with its content in different formats
type Article struct {
	Title           string
	ContentMarkdown string
	ContentHTML     string
}

// NewArticleManager creates a new ArticleManager instance with the specified base path.
func NewArticleManager(basePath string) *ArticleManager {
	return &ArticleManager{
		basePath: basePath,
	}
}

// Handle processes an uploaded file, validates it, and extracts its contents.
// It only accepts ZIP files containing markdown and images.
func (am *ArticleManager) Handle(file *multipart.FileHeader) error {
    am.article = Article{}

    if !am.isDirExist() {
        err := os.MkdirAll(am.basePath, 0755)
        if err != nil {
            return err
        }
    }

    title, extension, err := am.splitFile(file)
    if err != nil {
        return err
    }

    if extension != ".zip" {
        return errors.New("incorrect file extension: only ZIP files are supported")
    }

    srcFile, err := file.Open()
    if err != nil {
        return err
    }
    defer func() {
        closeErr := srcFile.Close()
        if closeErr != nil && err == nil {
            err = closeErr
        }
    }()

	tempZipPath := filepath.Join(am.basePath, file.Filename)
	
	// Create temporary file for the ZIP archive
	destFile, err := os.Create(tempZipPath)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := destFile.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Process the ZIP file using the original file name (without extension) as the directory name
	err = am.processZipFile(tempZipPath, title)

	// Clean up the temporary ZIP file
	os.Remove(tempZipPath)

	return err
}

// processZipFile extracts markdown and image files from a ZIP archive
// and creates the article structure in the file system.
func (am *ArticleManager) processZipFile(zipPath, zipTitle string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Create directory with the same name as the ZIP file
	articleDir := filepath.Join(am.basePath, zipTitle)
	err = os.MkdirAll(articleDir, 0755)
	if err != nil {
		return err
	}

	var mdFile *zip.File
	var mdFileName string

	// Find the first markdown file in the archive
	for _, file := range reader.File {
		if strings.HasSuffix(strings.ToLower(file.Name), ".md") {
			mdFile = file
			mdFileName = file.Name
			break
		}
	}

	if mdFile == nil {
		return errors.New("no markdown file found in the ZIP archive")
	}

	// Set article title from the markdown filename
	articleTitle := strings.TrimSuffix(filepath.Base(mdFileName), filepath.Ext(mdFileName))
	am.article.Title = articleTitle

	// Extract markdown content
	mdContent, err := am.extractFileContent(mdFile)
	if err != nil {
		return err
	}
	am.article.ContentMarkdown = mdContent

	// Save markdown content
	err = os.WriteFile(filepath.Join(articleDir, "content.md"), []byte(mdContent), 0644)
	if err != nil {
		return err
	}

	// Convert to HTML and save
	html, err := am.convertMarkdownToHTML(mdContent)
	if err != nil {
		return err
	}
	am.article.ContentHTML = html

	err = os.WriteFile(filepath.Join(articleDir, "content.html"), []byte(html), 0644)
	if err != nil {
		return err
	}

	// Create images directory and extract images
	imagesDir := filepath.Join(articleDir, "images")
	err = os.MkdirAll(imagesDir, 0755)
	if err != nil {
		return err
	}

	// Extract all image files
	for _, file := range reader.File {
		ext := strings.ToLower(filepath.Ext(file.Name))
		if am.isImageFile(ext) {
			err = am.extractFile(file, imagesDir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (am *ArticleManager) extractFileContent(file *zip.File) (string, error) {
	rc, err := file.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()

	content, err := io.ReadAll(rc)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (am *ArticleManager) extractFile(file *zip.File, destDir string) error {
	destPath := filepath.Join(destDir, filepath.Base(file.Name))

	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, rc)
	return err
}

// convertMarkdownToHTML converts markdown content to HTML using the markdown manager.
func (am *ArticleManager) convertMarkdownToHTML(markdown string) (string, error) {
	return manager.ArticleManage(markdown)
}

// isImageFile checks if a file extension corresponds to an image file.
func (am *ArticleManager) isImageFile(extension string) bool {
	imageExtensions := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, 
		".gif": true, ".bmp": true, ".svg": true, ".webp": true,
	}
	return imageExtensions[extension]
}

// CreateArticle initializes a new Article instance with optional title and content.
func (am *ArticleManager) CreateArticle(title string, content string) error {
	am.article = Article{
		Title:           title,
		ContentMarkdown: content,
	}
	
	if content != "" {
		html, err := am.convertMarkdownToHTML(content)
		if err != nil {
			return err
		}
		am.article.ContentHTML = html
	}
	
	return nil
}

// isDirExist checks if the base directory exists.
func (am *ArticleManager) isDirExist() bool {
	info, err := os.Stat(am.basePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

// splitFile extracts the name and extension from a file.
func (am *ArticleManager) splitFile(file *multipart.FileHeader) (string, string, error) {
	filename := file.Filename
	splited := strings.Split(filename, ".")
	if len(splited) < 2 {
		return "", "", errors.New("incorrect file name")
	}
	extension := "." + splited[len(splited)-1]
	name := strings.TrimSuffix(filename, extension)
	return name, extension, nil
}

// UpdateArticle updates an existing article with new content.
func (am *ArticleManager) UpdateArticle(title string, file *multipart.FileHeader) error {

	articlePath := filepath.Join(am.basePath, title)
	if !am.articleExists(title) {
		return errors.New("article does not exist")
	}
	
	err := os.RemoveAll(articlePath)
	if err != nil {
		return err
	}
	
	am.article.Title = title
	return am.Handle(file)
}

// DeleteArticle removes an article and all its files.
func (am *ArticleManager) DeleteArticle(title string) error {
	if !am.articleExists(title) {
		return errors.New("article does not exist")
	}
	
	return os.RemoveAll(filepath.Join(am.basePath, title))
}

// articleExists checks if an article with the given title exists.
func (am *ArticleManager) articleExists(title string) bool {
	info, err := os.Stat(filepath.Join(am.basePath, title))
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && info.IsDir()
}

