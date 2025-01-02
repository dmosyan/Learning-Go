package service

import (
	"testing"

	realdomain "github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/dto"
	"github.com/dmosyan/Learning-Go/apis/banking/errs"
	"github.com/dmosyan/Learning-Go/apis/banking/mocks/domain"
	"go.uber.org/mock/gomock"
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

func Test_should_return_error_from_server_if_account_cannot_be_created(t *testing.T) {
	// Arrange
	ctr := gomock.NewController(t)
	defer ctr.Finish()
	mockRepo := domain.NewMockAccountRepository(ctr)
	service := NewAccountService(mockRepo)

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000.00,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := service.NewAccount(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}

}
