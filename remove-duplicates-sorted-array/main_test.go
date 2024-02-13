package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestMain(t *testing.T) {
	type TestCase struct {
		name         string
		nums         []int
		expectedNums []int
		k            int
	}

	tests := []TestCase{
		// {
		// 	name:         "should return two unique elements",
		// 	nums:         []int{1, 1, 2},
		// 	expectedNums: []int{1, 2, 1},
		// 	k:            2,
		// },
		// {
		// 	name:         "should return two five elements",
		// 	nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		// 	expectedNums: []int{0, 1, 2, 3, 4, 2, 2, 3, 3, 4},
		// 	k: 5,
		// },
		{
			name:         "should return two five elements",
			nums:         []int{1, 2, 3},
			expectedNums: []int{1, 2, 3},
			k:            3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.nums)

			assert.Equal(t, tt.k, result)
			assert.DeepEqual(t, tt.expectedNums, tt.nums)
		})
	}
}
