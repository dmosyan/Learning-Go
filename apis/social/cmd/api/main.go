package main

import (
	"log"

	"github.com/dmosyan/Learning-Go/apis/social/internal/env"
	"github.com/dmosyan/Learning-Go/apis/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("PORT", ":3000"),
		db: dbConfig{
			addr:            env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable"),
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 10),
			maxConnLifetime: env.GetString("DB_MAX_CONN_LIFETIME", "15m"),
		},
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
