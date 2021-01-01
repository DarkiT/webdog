package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// read configuration from yaml
func ReadCfg() *Configuration {
	var cfg Configuration
	path := "./config.yml"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read configuration error:", err.Error())
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Println("configuration unmarshal error:", err.Error())
	}
	return &cfg
}
