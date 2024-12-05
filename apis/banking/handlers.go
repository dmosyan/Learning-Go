package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Banking App")
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
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

}
