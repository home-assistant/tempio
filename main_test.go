package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplatesJoin(t *testing.T) {
	dataJson := []byte(`{"list": ["hallo", "welt", "bla"]}`)
	dataTemplate := []byte(`{{ join ";" .list }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "hallo;welt;bla")
}

func TestTemplatesCalc(t *testing.T) {
	dataJson := []byte(`{"num": 3}`)
	dataTemplate := []byte(`{{ add .num 3 }}-{{ mul .num 3 }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "6-9")
}
