package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawl(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	appErr := req.Validate()

	// Assert
	if appErr.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Transaction type validation failed")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: DEPOSIT,
		Amount:          -100,
	}

	// Act
	appErr := req.Validate()

	// Assert
	if appErr.Message != "Amount cannot be less than zero" {
		t.Error("Amount validation failed")
	}

	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing amount")
	}
}
