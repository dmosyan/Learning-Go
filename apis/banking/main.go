package main

import "log"

const port = "3000"

func main() {

	registerRoutes()

	log.Println("starting banking service on port", port)
	startServer()

}
