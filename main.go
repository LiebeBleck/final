package main

import (
	"go1f/pkg/config"
	"go1f/pkg/db"
	"go1f/pkg/server"
	"log"
)

func main() {
	config.Init()

	if err := db.Init("scheduler.db"); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
