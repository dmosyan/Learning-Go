package domain

import (
	"strconv"

	"github.com/dmosyan/Learning-Go/apis/banking/errs"
	"github.com/dmosyan/Learning-Go/apis/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (cuomster_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while creating a new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last insert id for the new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
