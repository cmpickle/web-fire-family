package main

import (
	"./routes"
	"./app"
)

func main() {
	app.loadSettings()

	routes.InitRoutes()
}
