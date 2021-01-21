package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"text/template"
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

	// helper
	funcMap := template.FuncMap{
		"JoinString":   JoinString,
		"CalcAdd":      CalcAdd,
		"CalcSubtract": CalcSubtract,
		"CalcMultiply": CalcMultiply,
		"CalcDivide":   CalcDivide,
	}

	// generate template
	coreTemplate := template.New("tempIO").Funcs(funcMap)
	template.Must(coreTemplate.Parse(string(templateData)))

	// render
	err := coreTemplate.Execute(buf, *config)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
