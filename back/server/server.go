package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"plumlabs/back/api"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type Server struct {
	api api.API
	mux *http.ServeMux
	env Env
}

type Env struct {
	Port string
}

func NewServer(api api.API) *Server {

	s := Server{
		api: api,
		mux: http.NewServeMux(),
		env: *newEnv(),
	}

	return &s
}

func newEnv() *Env {
	return &Env{Port: loadPort()}
}

func (s *Server) StartWithGracefulShutdown() {
	srv := &http.Server{
		Addr:    ":" + s.env.Port,
		Handler: s.mux,
	}

	go func() {
		log.Printf("Server listening on port %s\n", s.env.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

func (s *Server) SetupRoutes() {
	s.mux.HandleFunc("/api/upload", s.api.ApiPostFile)
	s.mux.HandleFunc("/api/article/delete", s.api.ApiDeleteArticle)
	s.mux.HandleFunc("/api/articles/getall", s.api.ApiGetTitles)
	s.mux.HandleFunc("/api/article/get", s.api.ApiGetArticle) 

	fs := http.FileServer(http.Dir("./static"))
	s.mux.Handle("/", fs)
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
