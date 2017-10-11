package main

import (
	"net/http"

	"./routes"
)

func main() {
	http.ListenAndServe(":8080", routes.InitRoutes())
}
