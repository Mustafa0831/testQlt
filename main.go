package main

import (
	"log"
	"qlt/internal/server"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	paymentsRepository := server.NewPaymentsRepository()

	server := server.CreateApp(paymentsRepository)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

