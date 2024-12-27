package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomeRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{
		service: service.NewCustomerService(customerRepositoryDb),
	}
	ah := AccountHandler{
		service: service.NewAccountService(accountRepositoryDb),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}

func getDbClient() *sqlx.DB {
	dbUsr := os.Getenv("DB_USER")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSrc := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUsr, dbAddr, dbPort, dbName)
	c, err := sqlx.Open("mysql", dataSrc)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.SetConnMaxLifetime(time.Minute * 3)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	return c
}
