package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"yyds-pro/model"
)

//加载配置文件中的配置信息
func LoadConfig() (appConfig model.AppConfig, err error) {
	configFile, err := ioutil.ReadFile("./config/dev.yaml")
	if err != nil {
		log.Fatalf("yamlFile. Get err %v", err)
	}
	err = yaml.Unmarshal(configFile, &appConfig)
	return
}

//test-case用
func TestKafkaLoadConfig() (appConfig model.AppConfig, err error) {
	configFile, err := ioutil.ReadFile("../../config/dev.yaml")
	if err != nil {
		log.Fatalf("yamlFile. Get err %v", err)
	}
	err = yaml.Unmarshal(configFile, &appConfig)
	return
}
