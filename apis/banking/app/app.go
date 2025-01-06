package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/service"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/logger"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

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
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost).Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost).Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	addr := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// starting server
	logger.Info(fmt.Sprintf("starting OAuth server on %s:%s ...", addr, port))
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

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_ADDR",
		"DB_PORT",
	}
	emptyVars := []string{}

	for _, k := range envProps {
		if os.Getenv(k) == "" {
			emptyVars = append(emptyVars, k)
		}
	}
	if len(emptyVars) > 0 {
		logger.Error(fmt.Sprintf("environment variables %s not defined. Terminating application...", emptyVars))
	}
}
