package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplatesJoin(t *testing.T) {
	dataJSON := []byte(`{"list": ["hallo", "welt", "bla"]}`)
	dataTemplate := []byte(`{{ join ";" .list }}`)

	testConfig := readConfigBuffer(dataJSON)
	rendered := renderTemplateBuffer(testConfig, dataTemplate)

	assert.Equal(t, string(rendered), "hallo;welt;bla")
}

func TestTemplatesCalc(t *testing.T) {
	dataJSON := []byte(`{"num": 3}`)
	dataTemplate := []byte(`{{ add .num 3 }}-{{ mul .num 3 }}`)

	testConfig := readConfigBuffer(dataJSON)
	rendered := renderTemplateBuffer(testConfig, dataTemplate)

	assert.Equal(t, string(rendered), "6-9")
}
