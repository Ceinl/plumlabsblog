package server

import (
	"log"
	"net/http"
	"os"
	"plumlabs/back/api"

	"github.com/joho/godotenv"
)

type Server struct{
	api api.API
	mux *http.ServeMux
	port string
}

func NewServer(api api.API) *Server {

	s := Server{
    	api: api, 
    	mux: http.NewServeMux(),
    	port: loadPort(),
	}
	

	return &s
}

func (s *Server) Start() error {
	log.Printf("Server listening on port %s\n", s.port)
	return http.ListenAndServe(":"+s.port, s.mux)
}

func (s *Server) SetupRoutes(){
	s.mux.HandleFunc("", s.api.ApiPostFile)
	s.mux.HandleFunc("", s.api.ApiDeleteArticle)
	s.mux.HandleFunc("", s.api.ApiGetArticle)
	s.mux.HandleFunc("", s.api.ApiGetTitles)

}


func loadPort() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT not set in .env")
    }

    return port
}


