package main

import "fmt"

func main() {
	result := compress([]byte("aaabb"))
	fmt.Println(result)
}

// a a a b b
func compress(chars []byte) int {
	var s string
	count := 0
	current := ""

	for i, char := range chars {
		if string(char) != current {
			if count > 1 {
				s += fmt.Sprint(count)
			}

			s += string(char)
			current = string(char)

			count = 1

		} else {
			count++
		}

		if i == len(chars)-1 && count > 1 {
			s += fmt.Sprint(count)
		}
	}

	for i, value := range s {
		chars[i] = byte(value)
	}
	return len(s)
}
