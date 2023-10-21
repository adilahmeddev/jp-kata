package adapter

import (
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
	port   string
}

func NewServer(port string) *Server {
	return &Server{
		server: &http.Server{},
		port:   port,
	}

}

func (s *Server) ListenAndServe() error {
	s.server.Handler = &Handler{}
	s.server.Addr = ":" + s.port
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("listening and serving on %q\n", s.server.Addr)

	return nil
}
