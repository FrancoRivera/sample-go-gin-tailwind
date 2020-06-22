package main

import (
	"./config"
	"./internal"
)

func main() {
	internal.Initialize()
	config.SetupRoutes()
	internal.Router.Run(":8080")
}
