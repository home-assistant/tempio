package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

func JoinString(a []interface{}, delimiter string) string {
	var paramSlice []string
	for _, param := range a {
		paramSlice = append(paramSlice, param.(string))
	}
	return strings.Join(paramSlice, delimiter)
}

func renderTemplateFile(config *map[string]interface{}, file string) []byte {
	buf := &bytes.Buffer{}

	// read Template
	templateFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// helper
	funcMap := template.FuncMap{"JoinString": JoinString}

	// generate template
	coreTemplate := template.New("tempIO").Funcs(funcMap)
	template.Must(coreTemplate.Parse(string(templateFile)))

	// render
	err = coreTemplate.Execute(buf, *config)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
