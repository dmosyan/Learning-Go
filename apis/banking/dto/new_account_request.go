package dto

import (
	"strings"

	errs "github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000.0 {
		return errs.NewValidationError("to open an account you need to deposit at least $5000.00")
	}
	if strings.ToLower(r.AccountType) != "checking" && strings.ToLower(r.AccountType) != "saving" {
		return errs.NewValidationError("the account type should be checking or saving")
	}
	return nil
}
