package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_reverseWords(t *testing.T) {
	type TestCase struct {
		input  string
		output string
	}

	tests := []TestCase{
		{input: "Hello World", output: "World Hello"},
		{input: "Go is great", output: "great is Go"},
		{input: "  Leading and trailing spaces  ", output: "spaces trailing and Leading"},
		{input: "SingleWord", output: "SingleWord"},
		{input: "  ", output: ""},
		{input: "Multiple   spaces", output: "spaces Multiple"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := reverseWords(test.input)
			assert.Equal(t, test.output, result)
		})
	}
}
