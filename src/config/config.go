package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TypeAppConfig struct {
	Address string `json:"address"`
}

type TypeConfig struct {
	App TypeAppConfig `json:"app"`
}

var Config TypeConfig

func InitConfig() {
	configFilename := "default.json"

	configFile, err := ioutil.ReadFile("./config/" + configFilename)

	if err != nil {
		log.Println("config: read file error " + configFilename)
		log.Panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		log.Println("config: json unmarshal error " + configFilename)
		log.Panic(err)
	}

	log.Println("config: config " + configFilename + " loaded")
}
