package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dmosyan/Learning-Go/apis/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) customersHandler(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(customers)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(data)

}
