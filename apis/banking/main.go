package main

import (
	"github.com/dmosyan/Learning-Go/apis/banking/app"
)

func main() {

	mux := app.RegisterRoutes()

	app.StartServer(mux)

}
