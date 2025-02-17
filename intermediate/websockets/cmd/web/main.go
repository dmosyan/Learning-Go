package main

import (
	"log"
	"net/http"
)

func main() {

	routes := routes()

	log.Println("starting web server on :3000")

	_ = http.ListenAndServe(":3000", routes)

}
