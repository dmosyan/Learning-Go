package service

import (
	"testing"

	"github.com/dmosyan/Learning-Go/apis/banking/dto"
)

func Test_should_return_a_validation_error_response_when_the_request_not_valid(t *testing.T) {
	// Arrange
	req := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	// Act
	_, appErr := service.NewAccount(req)

	// Assert
	if appErr == nil {
		t.Error("Failed while testing a new account validation")
	}
}
