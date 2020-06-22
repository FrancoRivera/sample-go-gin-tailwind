package main

import (
	"./config"
	"./internal"
	"./repository"
)

func main() {
	internal.Initialize()
	config.SetupRoutes()
	repository.SetupDatabase()
	internal.Router.Run(":8080")
}
