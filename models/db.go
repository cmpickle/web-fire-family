package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Env struct {
	Db *sql.DB
}

var dbConnection string

func InitDB() (*sql.DB, error) {
	// Bootstrapping the setting

	//"fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family"
	dbConnection = "%v:%v@tcp(%v:)/%v"
	//Trying DB things here

	var err error
	//db, err = sql.Open("mysql", "fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
	NewDB("fireadmin:FireFamily@1@tcp(165.227.17.104:3306)/Fire_Family")
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
