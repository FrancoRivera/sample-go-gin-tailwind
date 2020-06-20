package main

import ("./internal"
	"./config")

func main() {
	internal.Initialize()
	config.SetupRoutes()
	internal.Router.Run(":8080")
}

