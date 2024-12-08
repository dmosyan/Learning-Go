package app

import "net/http"

func RegisterRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", greetHandler)
	mux.HandleFunc("/customers", customersHandler)

	return mux
}
