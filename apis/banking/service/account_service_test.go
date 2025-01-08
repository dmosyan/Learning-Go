package service

import (
	"testing"

	realdomain "github.com/dmosyan/Learning-Go/apis/banking/domain"
	"github.com/dmosyan/Learning-Go/apis/banking/dto"
	"github.com/dmosyan/Learning-Go/apis/banking/mocks/domain"
	errs "github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
	"go.uber.org/mock/gomock"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

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

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_error_from_server_if_account_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

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

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	accountWithId := account
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(account).Return(accountWithId, nil)

	// Act
	newAccount, appError := service.NewAccount(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Failed while mathching new account id")
	}
}
