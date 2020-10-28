package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func readConfigFile(file string) *map[string]interface{} {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
