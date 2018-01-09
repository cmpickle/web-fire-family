package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"

	"github.com/cmpickle/web-fire-family/app"
	"github.com/cmpickle/web-fire-family/models"
	"github.com/cmpickle/web-fire-family/routes"
)

var Dbdriver app.Dbdriver
var web app.Web

func main() {
	pwd, _ := os.Getwd()
	Dbdriver = Dbdriver.LoadSettings(pwd + "/configtest.yml")
	web = web.LoadSettings(pwd + "/configtest.yml")
	var addr string
	addr = ":" + strconv.Itoa(web.Port)
	db, err := models.InitDB(&Dbdriver)
	if err != nil {
		fmt.Errorf("Database wasn't initialized!")
	}

	env := models.Env{Db: db}

	http.ListenAndServe(addr, routes.InitRoutes(env))
}
