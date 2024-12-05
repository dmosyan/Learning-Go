package main

import "net/http"

func registerRoutes() {
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/goodbye", customersHandler)
}
