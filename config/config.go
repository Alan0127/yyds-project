package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"yyds-pro/model"
)

func LoadConfig() (appConfig model.AppConfig, err error) {
	configFile, err := ioutil.ReadFile("./config/dev.yaml")
	if err != nil {
		log.Fatalf("yamlFile. Get err %v", err)
	}
	err = yaml.Unmarshal(configFile, &appConfig)
	return
}
