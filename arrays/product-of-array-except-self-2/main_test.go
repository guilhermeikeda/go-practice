package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_productExceptSelf(t *testing.T) {
	type TestCase struct {
		input  []int
		output []int
	}

	tests := []TestCase{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{24, 12, 8, 6},
		},
	}
	for _, test := range tests {
		output := productExceptSelf(test.input)
		assert.DeepEqual(t, test.output, output)
	}
}
