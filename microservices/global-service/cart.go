package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type Cart struct {
	ID         int   `json:"id,omitempty"`
	CustomerID int   `json:"customerId,omitempty"`
	ProductIDs []int `json:"productIds,omitempty"`
}

var nextID int = 1
var carts = make([]Cart, 0)
var cartMux = http.NewServeMux()

func createShoppingCartService() *http.Server {

	cartMux.Handle("/carts", &validationMiddleware{next: http.HandlerFunc(cartsHandler)})

	//wrap mux with loggingMiddleware
	return &http.Server{
		Addr:    ":5000",
		Handler: &loggingMiddleware{next: cartMux},
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

type loggingMiddleware struct {
	next http.Handler
}

func (ln loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ln.next == nil {
		ln.next = cartMux
	}

	slog.Info(fmt.Sprintf("received %v request on route: %v", r.Method, r.URL.Path))
	now := time.Now()

	ln.next.ServeHTTP(w, r)

	slog.Info(fmt.Sprintf("response generated for %v request on route: %v. Duration: %v", r.Method, r.URL.Path, time.Since(now)))
}

type validationMiddleware struct {
	next http.Handler
}

func (val validationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if val.next == nil {
		log.Panic("no next handler defined for validationMiddleware")
	}

	if r.Method != http.MethodPost {
		val.next.ServeHTTP(w, r)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var c Cart
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := http.Head(fmt.Sprintf("http://localhost:3000/customers/%v", c.CustomerID))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if res.StatusCode == http.StatusNotFound {
		log.Print("Invalid customer ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := bytes.NewBuffer(data)
	r.Body = io.NopCloser(b)
	val.next.ServeHTTP(w, r)
}
