package main

import (
	"fmt"
	"slices"
)

func main() {
	output := reverseVowels("Icream") // acreim
	fmt.Println(output)
}

func reverseVowels(s string) string {
	i := 0
	j := len(s) - 1

	result := make([]byte, len(s))

	for i <= j {
		if isVowel(s[i]) && isVowel(s[j]) {
			result[i] = s[j]
			result[j] = s[i]
			i++
			j--
			continue
		}

		if !isVowel(s[i]) {
			result[i] = s[i]
			i++
		}

		if !isVowel(s[j]) {
			result[j] = s[j]
			j--
		}
	}

	return string(result)
}

func isVowel(s byte) bool {
	vowels := []byte{'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u'}

	return slices.Contains(vowels, s)
}
