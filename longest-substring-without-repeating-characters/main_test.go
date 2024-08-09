package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input := "dvdf"
	output := lengthOfLongestSubstring(input)

	assert.Equal(t, 3, output)
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	result := map[string]string{}
	start := 0

	for i := 0; i < len(s); i++ {
		current := string(s[i])
		if _, ok := result[current]; ok {
			result = map[string]string{}

			i = start
			start++
			continue
		}

		result[current] = current
		if len(result) > longest {
			longest = len(result)
		}
	}

	return longest
}

// Given a string s, find the length of the longest substring without repeating characters.
// func lengthOfLongestSubstring(s string) int {
// 	longest := ""

// 	for i := 0; i < len(s); i++ {
// 		start := i
// 		aux := ""
// 		for start < len(s) && !strings.Contains(aux, string(s[start])) {
// 			aux += string(s[start])
// 			start++
// 		}

// 		if len(longest) == 0 || len(longest) < len(aux) {
// 			longest = aux
// 		}
// 	}

// 	return len(longest)
// }
