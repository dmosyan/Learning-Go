package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Product struct {
	ID         int
	Name       string
	USDPerUnit float64
	Unit       string
}

type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}

func main() {
	cs := createCustomerService()
	ps := createProductService()
	scs := createShoppingCartService()

	go func() {
		cs.ListenAndServe()
	}()

	go func() {
		ps.ListenAndServe()
	}()

	go func() {
		scs.ListenAndServe()
	}()

	fmt.Println("services started, press < Enter > to shutdown")
	fmt.Scanln()
	shutDownServices(cs, ps, scs)
}

func createProductService() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(products)
		if err != nil {
			log.Print(err)
		}
		w.Header().Add("Content-type", "application/json")
		w.Write(data)
	})

	pattern := regexp.MustCompile(`^\/products\/(\d+)$`)
	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {

		matches := pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, p := range products {
			if id == p.ID {
				data, err := json.Marshal(p)
				if err != nil {
					log.Print(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.Header().Add("Content-type", "application/json")
				w.Write(data)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	})

	return &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

}

func createCustomerService() *http.Server {

	f, err := os.Open("../testdata/customers.csv")
	if err != nil {
		log.Fatal(err)
	}
	customers, err := readCustomers(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {

		data, err := json.Marshal(customers)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.Write(data)

	})

	pattern := regexp.MustCompile(`^\/customers\/(\d+?)$`)
	mux.HandleFunc("/customers/", func(w http.ResponseWriter, r *http.Request) {
		matches := pattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, c := range customers {
			if id == c.ID {
				data, err := json.Marshal(c)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Add("Content-type", "application/json")
				w.Write(data)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)

	})

	return &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
}

func readCustomers(r io.Reader) ([]Customer, error) {
	customers := make([]Customer, 0)
	csvReader := csv.NewReader(r)
	csvReader.Read() // throw away header
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			return customers, nil
		}
		if err != nil {
			return nil, err
		}
		var c Customer
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}
		c.ID = id
		c.FirstName = fields[1]
		c.LastName = fields[2]
		c.Address = fields[3]
		customers = append(customers, c)
	}
}

func shutDownServices(cs, ps, scs *http.Server) {
	cs.Shutdown(context.Background())
	ps.Shutdown(context.Background())
	scs.Shutdown(context.Background())
	fmt.Println("services stopped")
}