package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_kidsWithCandies(t *testing.T) {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)

	assert.DeepEqual(t, []bool{true, true, true, false, true}, result)
}
