package main

import (
	"fmt"
	"go-tickets/internel/config"
	"go-tickets/server"
)

func main() {
	//? Load environment variables
	cfg, err := config.LoadEnv()
	if err != nil {
		panic(fmt.Sprintf("failed to load environment variables: %v", err))
	}

	//? Connect to the database
	db := config.ConnectDatabase(cfg)

	//? Start the server
	server.StartServer(db, cfg)
}
