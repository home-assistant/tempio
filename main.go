package main

import (
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	configFile := flag.String("conf", "", "Config json file")
	templateFile := flag.String("template", "", "Template file")
	outFile := flag.String("out", "", "Output file")

	// parse command lines
	flag.Parse()

	// Get config
	config := readConfigFile(*configFile)

	// create & write corefile
	data := renderTemplateFile(config, *templateFile)
	err := ioutil.WriteFile(*outFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
