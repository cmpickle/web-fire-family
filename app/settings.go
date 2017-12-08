package app

import (
	"gopkg.in/yaml.v2"
	//"fmt"
	"log"
	//"io"
	"io/ioutil"
	//"os"
)

type Web struct {
	Port int `yaml:"webport,omitempty"`
}

type Dbdriver struct {
	Database string `yaml:"database,omitempty"`
	Driver   string `yaml:"driver,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Dbuser   string `yaml:"user,omitempty"`
	Dbpass   string `yaml:"pass,omitempty"`
	Port     int    `yaml:"dbport,omitempty"`
}

func (d Dbdriver) LoadSettingsDefault() Dbdriver {
	// slurping the config.yml file into memory.  and allowing the yaml framework handle the data read
	// This should get all setings from the file.
	// dat, err := ioutil.ReadFile("../config.yml")
	dat, err := ioutil.ReadFile("github.com/Xero67/web-fire-family/config.yml")
	yaml.Unmarshal(dat, &d)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return d
}

func (web Web) loadSettingsDefault() Web {
	// dat, err := ioutil.ReadFile("../config.yml")
	dat, err := ioutil.ReadFile("github.com/Xero67/web-fire-family/config.yml")
	yaml.Unmarshal(dat, &web)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return web
}

func (d Dbdriver) LoadSettings(s string) Dbdriver {
	// slurping the config.yml file into memory.  and allowing the yaml framework handle the data read
	// This should get all setings from the file.
	dat, err := ioutil.ReadFile(s)
	yaml.Unmarshal(dat, &d)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return d
}

func (web Web) LoadSettings(s string) Web {
	// slurping the config.yml file into memory.  and allowing the yaml framework handle the data read
	// This should get all setings from the file.
	dat, err := ioutil.ReadFile(s)
	yaml.Unmarshal(dat, &web)
	if err != nil {
		log.Fatal("cannot unmarshal data %v", err)
	}
	return web
}
