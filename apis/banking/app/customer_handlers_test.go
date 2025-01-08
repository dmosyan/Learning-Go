package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dmosyan/Learning-Go/apis/banking/dto"
	"github.com/dmosyan/Learning-Go/apis/banking/mocks/service"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {

	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)

	ch = CustomerHandlers{mockService}

	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {

	// Arrange
	tearDown := setup(t)
	defer tearDown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "John", City: "Anytown", Zipcode: "110011", DateofBirth: "2000-01-01", Status: "1"},
		{Id: "1002", Name: "Emily", City: "New York", Zipcode: "110012", DateofBirth: "2003-04-04", Status: "1"},
	}

	mockService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {

	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllCustomers("").Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
