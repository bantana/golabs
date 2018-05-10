package main

import (
	"log"
	"net/http"
	"os"
)

// Server ...
type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

// NewServer ...
func NewServer(options ...func(*Server)) *Server {
	s := &Server{
		logger: log.New(os.Stdout, "", 0),
		mux:    http.NewServeMux(),
	}

	for _, f := range options {
		f(s)
	}

	s.mux.HandleFunc("/", s.index)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	s.logger.Println("GET /")

	w.Write([]byte("Hello, World!"))
}

func main() {

}
