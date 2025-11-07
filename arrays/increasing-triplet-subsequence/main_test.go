package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_method(t *testing.T) {
	type TestCase struct {
		input  []int
		output bool
	}

	tests := []TestCase{
		// {
		// 	input:  []int{1, 2, 3, 4, 5},
		// 	output: true,
		// },
		// {
		// 	input:  []int{5, 4, 3, 2, 1},
		// 	output: false,
		// },
		// {
		// 	input:  []int{2, 1, 5, 0, 4, 6},
		// 	output: true,
		// },
		// {
		// 	input:  []int{20, 100, 10, 12, 5, 13},
		// 	output: true,
		// },
		{
			input:  []int{5, 4, 3, 2, 1},
			output: false,
		},
	}

	for _, test := range tests {
		output := increasingTriplet(test.input)
		assert.Equal(t, test.output, output)
	}
}
