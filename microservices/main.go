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
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := readCustomers()
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

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/servercontent", func(w http.ResponseWriter, r *http.Request) {
		customersFile, err := os.Open("./testdata/customers.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer customersFile.Close()
		// ideally we should check when the file was modified and pass that as time value
		// name value will not be used anyway (it checks the extension only)
		http.ServeContent(w, r, "customersdata.csv", time.Now(), customersFile)
	})

	http.HandleFunc("/servefile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=\"customers.csv\"")
		http.ServeFile(w, r, "./testdata/customers.csv")
	})

	http.HandleFunc("/service/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Service")
	})

	http.HandleFunc("/fprint", func(w http.ResponseWriter, r *http.Request) {
		customersFile, err := os.Open("./testdata/customers.csv")
		if err != nil {
			log.Fatal(err)
		}
		// defer customersFile.Close()
		// data, err := io.ReadAll(customersFile)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		//fmt.Fprint(w, string(data))

		// stream the data from the file to response
		io.Copy(w, customersFile)
	})

	var handlerFun http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	http.HandleFunc("/url/", handlerFun)

	http.Handle("/", myHandler("Customer Service"))

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	fmt.Println("Server started, press <Enter> to shutdown")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type myHandler string

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "energetic gophers")

	http.SetCookie(w, &http.Cookie{
		Name:    "session-id",
		Value:   "12345",
		Expires: time.Now().Add(24 * time.Hour * 365),
	})

	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintln(w, string(mh))
	fmt.Fprintln(w, r.Header)
}

type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
}

func readCustomers() ([]Customer, error) {
	f, err := os.Open("./testdata/customers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	customers := make([]Customer, 0)
	csvReader := csv.NewReader(f)
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
