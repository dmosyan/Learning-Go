package main

import (
	"github.com/dmosyan/Learning-Go/apis/banking/app"
)

func main() {

	router := app.RegisterRoutes()

	server := app.NewServer(":3000", router)

	app.StartServer(server)
}
