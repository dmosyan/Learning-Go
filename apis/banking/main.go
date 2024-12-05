package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Banking App")
	})

	s := http.Server{
		Addr: ":3000",
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
