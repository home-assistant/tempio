package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var config *map[string]interface{}
	configFile := flag.String("conf", "", "Config json file, can be omitted if used in a pipe")
	templateFile := flag.String("template", "", "Template file")
	outFile := flag.String("out", "", "Output file")

	// parse command lines
	flag.Parse()
	if *templateFile == "" {
		log.Fatal("Missing template argument")
	}
	if *outFile == "" {
		log.Fatal("Missing out argument")
	}

	// Get config
	config = readConfig(*configFile)

	// create & write corefile
	data := renderTemplateFile(config, *templateFile)
	err := ioutil.WriteFile(*outFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
