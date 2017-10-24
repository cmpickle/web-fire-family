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

type WebSettings struct {
	port int `port:"port",omitempty`
}

type database struct{
	// should move to the inline of the file. so we get whats under the mysql structure of the yaml file.
	Y string `yaml:"mysql,inline"`
}

type Dbdriver struct {
	Database database `yaml:"database,inline"`
	Driver string `yaml:"driver,omitempty"`
	Host string `yaml:"host,omitempty"`
	Dbuser string `yaml:"user,omitempty"`
	Dbpass string `yaml:"pass,omitempty"`
	Port int `port:"port,omitempty"`
}
func (d Dbdriver) LoadSettings() (Dbdriver) {
	// slurping the config.yml file into memory.  and allowing the yaml framework handle the data read
	// This should get all setings from the file.
	dat, err := ioutil.ReadFile("../config.yml")
	yaml.Unmarshal(dat, &d)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return d
}

func (web WebSettings) loadSettings() (WebSettings) {
	dat, err := ioutil.ReadFile("../config.yml")
	yaml.Unmarshal(dat,&web)
	if err != nil {
		log.Fatal("cannot unmarshal data %v" ,err)
	}
	return web
}
