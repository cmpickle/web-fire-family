package models

import (
	"database/sql"
	"log"

	"github.com/cmpickle/web-fire-family/app"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	Db *sql.DB
}

var dbConnection string

func InitDBdefault() (*sql.DB, error) {
	// Bootstrapping the setting

	//"fireadmin:FireFamily@1@tcp(50.63.80.1:3306)/Fire_Family"
	dbConnection = "%v:%v@tcp(%v:)/%v"
	//Trying DB things here

	var err error
	//db, err = sql.Open("mysql", "fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
	// NewDB("fisinvenory:P!Pawnshop1976@tcp(fisinventory.db.7722947.acb.hostedresource.net:3306)/fisinventory")
	NewDB("root:Pawnshop1976@localhost:3306)/fire_family")
	if err != nil {
		//error handling here
		log.Fatal("connection Error of %v", err)
		//fmt.Println("Conn")
		//fmt.Println(err)
	}
	if err = Db.Ping(); err != nil {
		//error handling here
		log.Fatal("No Ping of Database %v", err)
	}

	Db.SetMaxIdleConns(0)

	return Db, err
}
func InitDB(Dbdriver *app.Dbdriver) (*sql.DB, error) {
	// Bootstrapping the setting

	//"fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family"
	dbConnection = fmt.Sprintf("%v:%v@tcp(%v:)/%v", Dbdriver.Dbuser, Dbdriver.Dbpass, Dbdriver.Host, Dbdriver.Port, Dbdriver.Database)
	//Trying DB things here

	var err error
	//db, err = sql.Open("mysql", "fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
	// NewDB("fisinventory:P!Pawnshop1976@tcp(fisinventory.db.7722947.acb.hostedresource.net:3306)/fisinventory")
	NewDB("root:Pawnshop1976@tcp(127.0.0.1:3306)/fire_family")
	if err != nil {
		//error handling here
		log.Fatal("connection Error of %v", err)
		//fmt.Println("Conn")
		//fmt.Println(err)
	}
	if err = Db.Ping(); err != nil {
		//error handling here
		log.Fatal("No Ping of Database %v", err)
	}

	return Db, err
}

//SHOULD be a cross package global, isn't working, guess all db stuff is in routes now
var Db *sql.DB

func NewDB(dataSourceName string) (*sql.DB, error) {
	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		//error handling
	}
	if err = Db.Ping(); err != nil {
		//error handling
	}

	return Db, err
}
