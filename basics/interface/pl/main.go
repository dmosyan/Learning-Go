package main

import (
	"log"

	"github.com/dmosyan/Learning-Go/basics/interface/payment/payment"
)

type PaymentProcessor interface {
	ProcessPayment(amount float32) error
}

type Account interface {
	Available() float32
}

type PaymentMethod interface {
	PaymentProcessor
	Account
}

func main() {
	var pm PaymentMethod = payment.NewCreditCard("John Smith", "1234 5678 1234 5678", 12, 2022, 123, 1000)

	err := pm.ProcessPayment(500)
	if err != nil {
		log.Printf("Error processing payment: %v\n", err)
	} else {
		log.Printf("process payment successful, available balance: %v\n", pm.Available())
	}

	err = pm.ProcessPayment(600)
	if err != nil {
		log.Printf("Error processing payment: %v\n", err)
	} else {
		log.Printf("process payment successful, available balance: %v\n", pm.Available())
	}
}
