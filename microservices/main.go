package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.Handle("/", myHandler("Customer Service"))

	var handlerFun http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	http.HandleFunc("/url/", handlerFun)

	http.HandleFunc("/service/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Service")
	})

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
