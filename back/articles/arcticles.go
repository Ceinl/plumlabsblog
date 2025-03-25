package articles

import (
	"errors"
	"mime/multipart"
	"strings"
)

/*
	Створити унікальну папку для кожної статті.

	Зберегти md-файл у створеній папці.

	Викликати зовнішній пакет для конвертації md у html.

	Зберегти отриманий html-файл у відповідній папці.
*/

func HandleFile(files *multipart.FileHeader) error {

	name, extention, err:= splitName(files)
	if err != nil {
		return err
	}

	if isMarkDown(extention) {
		// Add to db create folder and a file	


	}else{
		return errors.New("incorrect file extention")
	}



	Savefile(name)
	return nil
}

func splitName(file *multipart.FileHeader) (string, string, error) {
	filename :=	file.Filename
	splited := strings.Split(filename, ".")
	if len(splited) != 2{
		return "", "", errors.New("incorrect file name")	
	}
	return splited[0] , splited[1], nil
}

func isMarkDown(extention string) bool {
	return extention == "md" 
}

func Savefile(name string){
	// TODO: Create folder with id and create in folder file with name of md file

}

