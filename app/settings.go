package app

import (
	"gopkg.in/yaml.v2"
	//"fmt"
	"log"
	//"io"
	"io/ioutil"
	//"os"
)

type database struct{
	y string `yaml:"mysql,inline"`
}

type driver struct {
	database database `yaml:"database,inline"`
	driver string `yaml:"driver,omitempty"`
	host string `yaml:"host,omitempty"`
	dbuser string `yaml:"user,omitempty"`
	dbpass string `yaml:"pass,omitempty"`
	port int `port:"port,omitempty"`
}
var dbdriver driver
func loadSettings() {
	dat, err := ioutil.ReadFile("../config.yml")
	yaml.Unmarshal(dat, &dbdriver)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
}
