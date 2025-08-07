package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	var words []string

	var word []string
	for i := len(s) - 1; i >= 0; i-- {
		if !isWhitespace(string(s[i])) {
			word = append(word, string(s[i]))
		}

		if (isWhitespace(string(s[i])) || i == 0) && len(word) > 0 {
			words = append(words, reverse(strings.Join(word, "")))
			word = []string{}
		}
	}

	return strings.Join(words, " ")
}

func isWhitespace(s string) bool {
	return s == " "
}

func reverse(s string) string {
	i := 0
	j := len(s) - 1

	result := make([]string, len(s))
	for i <= j {
		result[i] = string(s[j])
		result[j] = string(s[i])
		i++
		j--
	}

	return strings.Join(result, "")
}
