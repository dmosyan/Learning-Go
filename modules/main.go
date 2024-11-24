package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	rtr := mux.NewRouter()

	rtr.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to topics page"))
	})

	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to home page"))
	})

	err := http.ListenAndServe(":3000", rtr)
	if err != nil {
		log.Fatal(err)
	}

}
