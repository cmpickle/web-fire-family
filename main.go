package main

import (
	"net/http"

	"github.com/Xero67/web-fire-family/routes"
)

func main() {
	http.ListenAndServe(":8080", routes.InitRoutes())
}
