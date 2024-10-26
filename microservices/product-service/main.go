package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID         int
	Name       string
	USDPerUnit float64
	Unit       string
}

var products = []Product{
	{ID: 1, Name: "Apples", USDPerUnit: 1.99, Unit: "Pound"},
	{ID: 2, Name: "Oranges", USDPerUnit: 2.99, Unit: "Pound"},
	{ID: 3, Name: "Bread", USDPerUnit: 3.99, Unit: "Each"},
}

func main() {

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(products)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		//q := r.URL.Query().Get("id")

		parts := strings.Split(r.URL.Path, "/") // ["" "products" "3"]

		if len(parts) != 3 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, p := range products {
			if p.ID == id {
				data, err := json.Marshal(p)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-Type", "application/json")
				w.Write(data)
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})

	s := http.Server{
		Addr: ":4000",
	}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}
