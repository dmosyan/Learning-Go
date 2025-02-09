package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmosyan/Learning-Go/apis/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type dbConfig struct {
	addr            string
	maxOpenConns    int
	maxIdleConns    int
	maxConnLifetime string
}

type config struct {
	addr string
	db   dbConfig
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthCheckHandler)
	})

	return r
}

func (a *application) run(mux http.Handler) error {

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
