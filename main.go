package main

import (
	"net/http"
	//"database/sql"
	//"./models"
	"./routes"
)

/*//Dependency injection? Screw it, going with a global
type Env struct {
	db *sql.DB
}*/

func main() {
	//Ignore this stuff, was trying to figure out dependency injection and global variables
	/*db, err := models.NewDB("fireadmin:FireFamily@1@165.227.17.104:3306/FireFamilyDB")
	if err != nil {
		//Error handling here
	}
	defer db.Close()
	env := &Env{db: db}


	//Creates the global variable, also good reason for config file here
	models.NewDB("fireadmin:FireFamily@1@165.227.17.104:3306/FireFamilyDB")
	*/
	http.ListenAndServe(":8000", routes.InitRoutes())
}
