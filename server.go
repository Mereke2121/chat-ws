package app

import (
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.Server = &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.ListenAndServe()
}
