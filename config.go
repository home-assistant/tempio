package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func readConfig(file string) *map[string]interface{} {
	if file == "" {
		return readConfigPipe()
	} else {
		return readConfigFile(file)
	}
}

func readConfigPipe() *map[string]interface{} {
	var config *map[string]interface{}
	defer os.Stdin.Close()
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		err := json.NewDecoder(os.Stdin).Decode(&config)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer os.Stdout.Close()
	return config
}

func readConfigFile(file string) *map[string]interface{} {
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// Parse json
	var config map[string]interface{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
