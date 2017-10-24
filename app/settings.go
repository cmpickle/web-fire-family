package app

import (
	"gopkg.in/yaml.v2"
	//"fmt"
	"log"
	//"io"
	"io/ioutil"
	//"os"
)

type web  struct {
	y string `yaml:"web,inline"`
}

type webSettings struct {
	port int `port:"port",omitempty`
}

type database struct{
	// should move to the inline of the file. so we get whats under the mysql structure of the yaml file.
	y string `yaml:"mysql,inline"`
}

type dbdriver struct {
	database database `yaml:"database,inline"`
	driver string `yaml:"driver,omitempty"`
	host string `yaml:"host,omitempty"`
	dbuser string `yaml:"user,omitempty"`
	dbpass string `yaml:"pass,omitempty"`
	port int `port:"port,omitempty"`
}
func (d dbdriver) loadSettings() (dbdriver) {
	// slurping the config.yml file into memory.  and allowing the yaml framework handle the data read
	// This should get all setings from the file.
	dat, err := ioutil.ReadFile("../config.yml")
	yaml.Unmarshal(dat, &d)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return d
}

func (web webSettings) loadSettings() (webSettings) {
	dat, err := ioutil.ReadFile("../config.yml")
	yaml.Unmarshal(dat,&web)
	if err != nil {
		log.Fatal("cannot unmarshal data %v" ,err)
	}
	return web
}
