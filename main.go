package main

import (
	"bank/application"
	"bank/infrastructure"
	"bank/presentation"
)

func main() {
	db := infrastructure.NewDatabase()
	accountService := application.NewAccountService(db)

	presentation.StartHTTPServer(accountService)
}
