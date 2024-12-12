package main

import (
	"github.com/dmosyan/Learning-Go/apis/banking/app"
	"github.com/dmosyan/Learning-Go/apis/banking/logger"
)

func main() {

	router := app.RegisterRoutes()

	server := app.NewServer(":3000", router)

	logger.Info("starting the application...")
	app.StartServer(server)
}
