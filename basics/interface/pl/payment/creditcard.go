package payment

import "errors"

type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

func NewCreditCard(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int, availableCredit float32) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
		availableCredit: availableCredit,
	}
}

func (c *CreditCard) Available() float32 {
	return c.availableCredit
}

func (c *CreditCard) ProcessPayment(amount float32) error {
	if c.availableCredit < amount {
		return errors.New("insufficient funds to complete payment")
	}

	c.availableCredit -= amount
	return nil
}
