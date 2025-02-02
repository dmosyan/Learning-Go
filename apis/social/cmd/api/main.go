package main

import (
	"log"

	"github.com/dmosyan/Learning-Go/apis/social/internal/env"
	"github.com/dmosyan/Learning-Go/apis/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetEnv("PORT", ":3000"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
