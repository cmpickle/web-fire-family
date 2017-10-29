package main

import (
	"fmt"
	"net/http"

	"./models"
	"./routes"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		fmt.Errorf("Database wasn't initialized!")
	}

	env := models.Env{Db: db}

	http.ListenAndServe(":8000", routes.InitRoutes(env))
}
