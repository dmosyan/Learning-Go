package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/service/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Service")
	})

	http.HandleFunc("/fprint/", func(w http.ResponseWriter, r *http.Request) {
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
