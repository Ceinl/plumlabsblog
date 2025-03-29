package articles

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func HandleFile(files *multipart.FileHeader) error {

	name, extention, err:= splitName(files)
	if err != nil {
		return err
	}

	if extention != "md"{
		return errors.New("incorrect file extention")
	}

	err = createFile(name, files)
	if err != nil{
		return err
	}

	return nil
}
func isDirExist(name string) bool{
	dirpath := filepath.Join("storage", "articles", name)
	info, err := os.Stat(dirpath)

	if os.IsNotExist(err){
		return false 
	}

	return err == nil && info.IsDir() 
}

func splitName(file *multipart.FileHeader) (string, string, error) {
	filename :=	file.Filename
	splited := strings.Split(filename, ".")
	if len(splited) != 2{
		return "", "", errors.New("incorrect file name")	
	}
	return splited[0] , splited[1], nil
}

func createFile(name string, fh *multipart.FileHeader) error {
	dirpath := filepath.Join("storage", "articles", name)
	
	err := os.MkdirAll(dirpath, 0755)
	if err != nil {
		return err
	}

	file, err := fh.Open()	
	if err != nil{
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil{
		return err
	}

	out, err := os.Create(filepath.Join(dirpath, "content" + ".md"))
	if err != nil {
		return err
	}
	defer out.Close()

	out.Write(data)

	return nil
}

