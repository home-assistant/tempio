package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinString(t *testing.T) {
	dataJson := []byte(`{"list": ["hallo", "welt", "bla"]}`)
	dataTemplate := []byte(`{{ JoinString .list ";" }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "hallo;welt;bla")
}

func TestCalcAdd(t *testing.T) {
	dataJson := []byte(`{"num": 1}`)
	dataTemplate := []byte(`{{ CalcAdd .num 2 }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "3")
}

func TestCalcSubtract(t *testing.T) {
	dataJson := []byte(`{"num": 3}`)
	dataTemplate := []byte(`{{ CalcSubtract .num 2 }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "1")
}

func TestCalcMultiply(t *testing.T) {
	dataJson := []byte(`{"num": 2}`)
	dataTemplate := []byte(`{{ CalcMultiply .num 3 }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "6")
}

func TestCalcDivide(t *testing.T) {
	dataJson := []byte(`{"num": 6}`)
	dataTemplate := []byte(`{{ CalcDivide .num 2 }}`)

	test_config := readConfigBuffer(dataJson)
	rendered := renderTemplateBuffer(test_config, dataTemplate)

	assert.Equal(t, string(rendered), "3")
}
