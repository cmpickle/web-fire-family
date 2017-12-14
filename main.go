package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cmpickle/web-fire-family/app"
	"github.com/cmpickle/web-fire-family/models"
	"github.com/cmpickle/web-fire-family/routes"
)

var Dbdriver app.Dbdriver
var web app.Web

func main() {
	Dbdriver = Dbdriver.LoadSettings("github.com/cmpickle/web-fire-family/config.yml")
	web = web.LoadSettings("github.com/cmpickle/web-fire-family/config.yml")
	var addr string
	addr = ":" + strconv.Itoa(web.Port)
	db, err := models.InitDB(&Dbdriver)
	if err != nil {
		fmt.Errorf("Database wasn't initialized!")
	}

	env := models.Env{Db: db}

	http.ListenAndServe(addr, routes.InitRoutes(env))
}
