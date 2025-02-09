package main

import (
	"log"

	"github.com/dmosyan/Learning-Go/apis/social/internal/db"
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

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxConnLifetime)
	if err != nil {
		log.Panic(err)
	}

	db.Close()
	log.Println("db connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
