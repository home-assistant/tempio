package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var TempioVersion string

func main() {
	var config *map[string]interface{}
	configFile := flag.String("conf", "", "Config json file, can be omitted if used in a pipe")
	templateFile := flag.String("template", "", "Template file")
	outFile := flag.String("out", "", "Output file, if not defined output will be to console")

	flag.Usage = func() {
		writer := flag.CommandLine.Output()

		fmt.Fprintf(writer, "Version: %s\n", TempioVersion)
		fmt.Fprintf(writer, "Documentation: https://github.com/home-assistant/tempio\n\n")

		flag.PrintDefaults()
	}

	// parse command lines
	flag.Parse()
	if *templateFile == "" {
		log.Fatal("Missing template argument")
	}

	// Get config
	config = readConfig(*configFile)

	// create & write corefile
	data := renderTemplateFile(config, *templateFile)
	if *outFile == "" {
		fmt.Println(string(data))
	} else {
		err := os.WriteFile(*outFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
