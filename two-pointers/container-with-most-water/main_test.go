package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_method(t *testing.T) {
	type TestCase struct {
		input  []int
		output int
	}

	tests := []TestCase{
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
	}

	for _, test := range tests {
		output := maxArea(test.input)
		assert.Equal(t, test.output, output)
	}
}
