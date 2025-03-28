package articles

import (
	"errors"
	"mime/multipart"
	"plumlabs/back/storage"
	"strings"
)

/*
	створити унікальну папку для кожної статті.
	зберегти md-файл у створеній папці.
	викликати зовнішній пакет для конвертації md у html.
	зберегти отриманий html-файл у відповідній папці.
*/

func HandleFile(files *multipart.FileHeader) error {

	name, extention, err:= splitName(files)
	if err != nil {
		return err
	}

	if extention == "md" {
		db, err := storage.Open()
		
		if err != nil {
			return err
		}

		storage.Init(db)
		storage.InsertTable(db,name)
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

func Create(file multipart.File, name string){
	
}

func Savefile(name string){
	// TODO: Create folder with id and create in folder file with name of md file
}

