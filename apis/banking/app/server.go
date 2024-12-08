package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

const port = "3000"

func StartServer(mux *http.ServeMux) {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	log.Println("starting banking service on port", port)
	log.Print("server started, press <Enter> to shutdown")
	fmt.Scanln()
	err := s.Shutdown(context.Background())
	if err != nil {
		log.Print("failed to stop the server\n")
	}
	log.Print("server is stopped\n")
}
