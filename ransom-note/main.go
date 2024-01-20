package main

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/
func canConstruct(ransomNote string, magazine string) bool {

	inputMap := make(map[rune]int)
	for _, letter := range ransomNote {
		inputMap[letter]++
	}

	for _, letter := range magazine {
		if _, ok := inputMap[letter]; ok {
			inputMap[letter]--
		}
	}

	for _, value := range inputMap {
		if value > 0 {
			return false
		}
	}

	return true
}
