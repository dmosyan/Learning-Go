package domain

type Transaction struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == "withdrawal" {
		return true
	}
	return false
}
