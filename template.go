package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/Masterminds/sprig"
)

func renderTemplateFile(config *map[string]interface{}, file string) []byte {
	// read Template
	templateFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cant read file %s - %s", file, err)
	}

	return renderTemplateBuffer(config, templateFile)
}

func renderTemplateBuffer(config *map[string]interface{}, templateData []byte) []byte {
	buf := &bytes.Buffer{}

	// generate template
	coreTemplate := template.New("tempIO").Funcs(sprig.TxtFuncMap())
	template.Must(coreTemplate.Parse(string(templateData)))

	// render
	err := coreTemplate.Execute(buf, *config)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
