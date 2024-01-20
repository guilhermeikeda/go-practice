package main

import (
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestMajorityElements(t *testing.T) {
	type TestCase struct {
		input    []int
		expected int
	}

	tests := []TestCase{
		{
			input:    []int{3, 2, 3},
			expected: 3,
		},
		{
			input:    []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			input:    []int{6, 5, 5},
			expected: 5,
		},
		{
			input:    []int{8, 8, 7, 7, 7},
			expected: 7,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(tt *testing.T) {
			result := majorityElements(test.input)
			assert.Equal(t, test.expected, result)
		})
	}

}
