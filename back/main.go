package main

import (
	"log"
	"net/http"
	"plumlabs/back/articles"
)
func main() {
	// Admin route to upload article to server
	http.HandleFunc("/upload", uploadFile) 
	
	// p + l alphabet order
	server(":1612") 
}


func server(port string) {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/",fs)

	log.Println("Server started at " + port)
	http.ListenAndServe(port,nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("UploadFile called")
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "File too big", http.StatusBadRequest)
		log.Println("Error parsing form:", err)
		return
	}

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}

	// Get the first file from the form data
	file := files[0]
	log.Println("Uploading file:", file.Filename)

	// Open the file
	fileHandle, err := file.Open()
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		log.Println("Error opening file:", err)
		return
	}
	am := articles.NewArticleManager("")
	am.Handle()
	defer fileHandle.Close()

}
