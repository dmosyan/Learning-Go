package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	receiveOrders()
	fmt.Println(orders)

}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 23, "status": 1}`,
	`{"productCode": 3333, "quantity": 7.52, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

func receiveOrders() {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}

		orders = append(orders, newOrder)
	}
}
