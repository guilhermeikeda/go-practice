package main

import "fmt"

func main() {
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var result string
	var remain string

	var minLength int

	if len(word2) > len(word1) {
		minLength = len(word1)
		remain = string(word2[len(word1):])
	} else {
		minLength = len(word2)
		remain = string(word1[len(word2):])
	}

	for i := 0; i < minLength; i++ {
		result += string(word1[i])
		result += string(word2[i])
	}

	result += remain

	return result
}
