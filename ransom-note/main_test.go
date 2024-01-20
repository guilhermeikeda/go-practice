package main

import (
	"strconv"
	"testing"

	"gotest.tools/assert"
)

func TestCanConstruct(t *testing.T) {
	type TestCase struct {
		ransomNote string
		magazine   string

		expected bool
	}
	tests := []TestCase{
		{
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			expected:   true,
		},
		{
			ransomNote: "bg",
			magazine:   "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",
			expected:   true,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(tt *testing.T) {
			result := canConstruct(test.ransomNote, test.magazine)
			assert.Equal(t, test.expected, result)
		})
	}
}
