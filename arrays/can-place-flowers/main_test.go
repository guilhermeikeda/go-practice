package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canPlaceFlowers(t *testing.T) {
	type TestCase struct {
		flowerbed []int
		n         int
		output    bool
	}

	tests := []TestCase{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			output:    true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			output:    false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			output:    false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
			n:         2,
			output:    true,
		},
		{
			flowerbed: []int{0, 0, 1, 0, 0},
			n:         1,
			output:    true,
		},
	}

	for _, test := range tests {
		result := canPlaceFlowers(test.flowerbed, test.n)

		assert.Equal(t, test.output, result)
	}
}
