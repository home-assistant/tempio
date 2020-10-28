package main

import (
	"bytes"
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

func renderTemplateFile(config *map[string]interface{}, coreFile string) []byte {
	buf := &bytes.Buffer{}

	// helper
	funcMap := template.FuncMap{"JoinString": JoinString}

	// generate template
	coreTemplate := template.New("tempio").Funcs(funcMap)
	template.Must(coreTemplate.ParseFiles(coreFile))

	// render
	err := coreTemplate.Execute(buf, *config)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
