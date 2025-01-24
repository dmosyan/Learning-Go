package main

import (
	"log"

	"github.com/dmosyan/Learning-Go/apis/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetEnv("PORT", ":3000"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
