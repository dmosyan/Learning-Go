package main

import (
	"github.com/dmosyan/Learning-Go/apis/banking/app"
	"github.com/dmosyan/Learning-Go/apis/shared/pkg/banking-lib/logger"
)

func main() {

	logger.Info("starting the application...")
	app.Start()
}
