package main

import (
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestTwoSum(t *testing.T) {
	type TestCase struct {
		input  []int
		target int

		expected []int
	}

	tests := []TestCase{
		{
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			input:    []int{3, 2, 3},
			target:   6,
			expected: []int{0, 2},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(tt *testing.T) {
			result := twoSum(test.input, test.target)
			assert.DeepEqual(t, test.expected, result)
		})
	}
}
