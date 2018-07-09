package app

import (
	"fmt"
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
func (s *Server) Init(port string) {
	log.Println("Initializing server...")
	s.port = fmt.Sprintf(":%s", port)
}

// Start iniciar o servidor
func (s *Server) Start() {
	log.Println("Starting server on port", s.port)
	http.ListenAndServe(s.port, nil)
}
