package main

import (
	"fmt"
	"net/http"
	"./app"
	"./models"
	"./routes"
	"strconv"
)
var Dbdriver app.Dbdriver
var web app.Web
func main() {
	Dbdriver = Dbdriver.LoadSettings("./config.yml")
	web = web.LoadSettings("./config.yml")
	var addr string
	addr = ":" + strconv.Itoa(web.Port)
	db, err := models.InitDB(&Dbdriver)
	if err != nil {
		fmt.Errorf("Database wasn't initialized!")
	}

	env := models.Env{Db: db}

	http.ListenAndServe(addr, routes.InitRoutes(env))
}
