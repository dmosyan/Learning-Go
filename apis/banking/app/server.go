package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func NewServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: handler,
	}
}

func StartServer(s *http.Server) {

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	log.Print("server started, press <Enter> to shutdown")
	fmt.Scanln()
	err := s.Shutdown(context.Background())
	if err != nil {
		log.Print("failed to stop the server\n")
	}
	log.Print("server is stopped\n")
}
