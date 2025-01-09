package main

import "log"

func main() {
	cfg := config{
		addr: ":3005",
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
