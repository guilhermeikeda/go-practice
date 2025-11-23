package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_method(t *testing.T) {
	type TestCase struct {
		nums   []int
		k      int
		output int
	}

	tests := []TestCase{
		{
			nums:   []int{2, 2, 2, 3, 1, 1, 4, 1},
			k:      4,
			output: 2,
		},
	}

	for _, test := range tests {
		output := maxOperations(test.nums, test.k)
		assert.Equal(t, test.output, output)
	}
}
