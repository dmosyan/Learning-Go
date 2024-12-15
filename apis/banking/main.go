package main

import (
	"github.com/dmosyan/Learning-Go/apis/banking/app"
	"github.com/dmosyan/Learning-Go/apis/banking/logger"
)

func main() {

	logger.Info("starting the application...")
	app.Start()
}
