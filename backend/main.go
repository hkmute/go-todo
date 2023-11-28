package main

import (
	"go-todo/packages/config"
	"go-todo/packages/db"
	"go-todo/packages/server"
)

func init() {
	config.LoadEnv()
}

func main() {
	db.DB.Connect()
	defer db.DB.Close()

	server.Run()
}
