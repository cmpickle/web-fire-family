package main

import (
	"fmt"
	"net/http"
	"strconv"

	"./app"
	"./models"
	"./routes"
	// "github.com/Xero67/web-fire-family/app"
	// "github.com/Xero67/web-fire-family/models"
	// "github.com/Xero67/web-fire-family/routes"
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
