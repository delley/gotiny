package app

import (
	"log"
	"net/http"
)

// Server eh uma struct que representa os dados do servidor
type Server struct {
	port string
}

// NewServer cria um novo servidor
func NewServer() Server {
	return Server{}
}

// Init inicializa a configuracao do servidor
func (s *Server) Init() {
	log.Println("Initializing server...")
	s.port = ":8080"
}

// Start iniciar o servidor
func (s *Server) Start() {
	log.Println("Starting server...")
	http.ListenAndServe(s.port, nil)
}
