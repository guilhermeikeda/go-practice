package main

import (
	"fmt"
	"slices"
)

func main() {

}

func reverseVowels(s string) string {
	vowels := getVowels(s)
	vowels = reverse(vowels)
	vowelsIndex := 0

	var result string
	for _, value := range s {
		if !isVowel(value) {
			result = fmt.Sprintf("%s%s", result, string(value))
			continue
		}

		result = fmt.Sprintf("%s%s", result, string(vowels[vowelsIndex]))
		vowelsIndex++
	}

	return result
}

func getVowels(input string) []rune {
	result := []rune{}
	for _, char := range input {
		if isVowel(char) {
			result = append(result, char)
		}
	}

	return result
}

func reverse(input []rune) []rune {
	var result []rune

	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}

	return result
}

func isVowel(s rune) bool {
	vowels := []rune{'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u'}

	return slices.Contains(vowels, s)
}
