package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := longestCommonPrefix(
		[]int{25, 43, 123, 1234},
		[]int{1234, 2, 5657})
	fmt.Println(result)
}

func longestCommonPrefix(first []int, second []int) int {
	var result int

	for _, first := range first {
		firstString := strconv.Itoa(first)
		for i := 1; i <= len(firstString); i++ {
			prefix := firstString[0:i]

			if result > len(prefix) {
				continue
			}

			for _, value := range second {
				valueString := strconv.Itoa(value)
				if strings.HasPrefix(valueString, prefix) {
					if len(prefix) > result {
						result = len(prefix)
						break
					}
				}
			}
		}
	}

	return result
}
