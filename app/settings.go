package app

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"log"
)

type database struct{
	y string `mysql:",inline"`
}

type driver struct {
	database `mysql:",inline"`
	driver string `driver:"a,omitempty"`
	host string `host:"a,omitempty"`
	dbuser string `dbuser:"a,omitempty"`
	dbpass string `dbpass:"a,omitempty`
	port int `port:a,omitempty`
}



