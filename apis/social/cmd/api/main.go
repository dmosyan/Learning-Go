package main

import (
	"github.com/dmosyan/Learning-go/apis/social/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetEnv("PORT", "8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
