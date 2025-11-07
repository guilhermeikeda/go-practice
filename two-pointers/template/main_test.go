package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_method(t *testing.T) {
	type TestCase struct {
		input  string
		output string
	}

	tests := []TestCase{
		{
			input:  "",
			output: "",
		},
	}

	for _, test := range tests {
		output := method()
		assert.Equal(t, test.output, output)
	}
}
