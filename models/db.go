package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//SHOULD be a cross package global, isn't working, guess all db stuff is in routes now
var Db *sql.DB

func NewDB(dataSourceName string) {
	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		//error handling
	}
	if err = Db.Ping(); err != nil {
		//error handling
	}
}