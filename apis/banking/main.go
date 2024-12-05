package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Banking App")
	})

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := getCustomers()
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(customers)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(data)

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
