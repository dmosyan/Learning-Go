package app

import (
	"net/http"

	"github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/service"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	mux := mux.NewRouter()

	// ch := CustomerHandlers{
	// 	service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	// }
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomeRepositoryDb()),
	}

	mux.HandleFunc("/customers", ch.customersHandler).Methods(http.MethodGet)

	return mux
}
