package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("aaaaaa", "bbaaaa"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	pointerT := 0
	pointerS := 0

	for pointerT < len(t) {
		if s[pointerS] == t[pointerT] {
			pointerS++

			if pointerS == len(s) {
				return true
			}
		}

		pointerT++
	}

	return false
}
