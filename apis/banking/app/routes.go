package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	mux := mux.NewRouter()

	mux.HandleFunc("/", greetHandler)
	mux.HandleFunc("/customers", customersHandler)

	return mux
}
