package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Cart struct {
	ID         int   `json:"id,omitempty"`
	CustomerID int   `json:"customerId,omitempty"`
	ProductIDs []int `json:"productIds,omitempty"`
}

var nextID int = 1
var carts = make([]Cart, 0)

func createShoppingCartService() *http.Server {
	var cartMux = http.NewServeMux()

	cartMux.HandleFunc("/carts", cartsHandler)

	//wrap mux with loggingMiddleware
	return &http.Server{
		Addr:    ":5000",
		Handler: cartMux,
	}
}

func cartsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := json.Marshal(carts)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.Write(data)

	case http.MethodPost:
		var c Cart
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&c)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		c.ID = nextID
		nextID++
		carts = append(carts, c)
		data, err := json.Marshal(carts)
		if err != nil {
			log.Print(err)
			fmt.Fprintf(w, "failed to return created cart data")
			return
		}
		w.Write(data)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
