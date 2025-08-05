package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseVowels(t *testing.T) {
	type TestCase struct {
		input  string
		output string
	}

	tests := []TestCase{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
		{"aA", "Aa"},
		{"", ""},
		{"bcd", "bcd"},
	}

	for _, tc := range tests {
		result := reverseVowels(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
