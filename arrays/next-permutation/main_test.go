package nextpermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPivot(t *testing.T) {
	pivot := findPivot([]int{2, 4, 1, 7, 5, 0})
	assert.Equal(t, 1, *pivot)
}

func TestReverse(t *testing.T) {
	got := reverse([]int{1, 2, 3})
	assert.Equal(t, []int{3, 2, 1}, got)

	got = reverse([]int{1, 2, 3, 4})
	assert.Equal(t, []int{4, 3, 2, 1}, got)
}

func TestFindSuccessor(t *testing.T) {
	got := findSuccessor(1, []int{2, 4, 1, 7, 5, 0})
	assert.Equal(t, 5, *got)
}

func TestNextPermutation(t *testing.T) {
	type Test struct {
		input  []int
		output []int
	}

	tests := []Test{
		// {
		// 	input:  []int{1, 2, 3},
		// 	output: []int{1, 2, 3},
		// },
		{
			input:  []int{2, 4, 1, 7, 5, 0},
			output: []int{2, 4, 5, 0, 1, 7},
		},
	}

	for _, test := range tests {
		got := nextPermutation(test.input)

		assert.Equal(t, test.output, got)
	}
}
