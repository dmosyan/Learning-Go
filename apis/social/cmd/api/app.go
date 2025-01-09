package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) run() error {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Printf("starting server on %s", a.config.addr)

	return srv.ListenAndServe()
}
