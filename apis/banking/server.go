package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func startServer() {
	s := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
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
