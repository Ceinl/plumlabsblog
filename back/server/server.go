package server

import (
	"net/http"
	"plumlabs/back/api"
)

type Server struct{
	api api.API
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{}
}

