package main

import (
	"net/http"

	"./routes"
)

func main() {
	http.ListenAndServe(":8000", routes.InitRoutes())
}
