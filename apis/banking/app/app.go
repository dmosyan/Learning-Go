package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/service"
	"github.com/gorilla/mux"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("server env variable not defined")
	}
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("db env variables are not defined")
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	// ch := CustomerHandlers{
	// 	service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	// }
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomeRepositoryDb()),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}
