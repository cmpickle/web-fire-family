package app

import "testing"

func SettingTest(t testing.T){

	//Arrange
	var driver string
	var host string
	var dbpass string
	var dbuser string
	var port int

	driver = "mysql"
	host = "165.227.17.104"
	dbuser = "fireadmin"
	dbpass = "FireFamily@1"
	port = 3306

	// read the yaml and verify that it matchs the settings above.
	db := database{driver:"mysql", }
}
