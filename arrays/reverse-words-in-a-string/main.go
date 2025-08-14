package main

func main() {

}

func reverseWords(s string) string {
	// Convert to []byte for in-place mutation
	b := []byte(s)
	b = cleanSpaces(b)      // Step 1: remove extra spaces
	reverse(b, 0, len(b)-1) // Step 2: reverse the entire slice

	// Step 3: reverse each word
	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			start = i + 1
		}
	}
	return string(b)
}

// Reverse helper
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// Remove leading/trailing/multiple spaces
func cleanSpaces(b []byte) []byte {
	n := len(b)
	write := 0
	read := 0

	// Skip leading spaces
	for read < n && b[read] == ' ' {
		read++
	}

	for read < n {
		// Copy non-space chars
		if b[read] != ' ' {
			b[write] = b[read]
			write++
		} else if write > 0 && b[write-1] != ' ' {
			// Keep one space between words
			b[write] = ' '
			write++
		}
		read++
	}

	// Remove trailing space
	if write > 0 && b[write-1] == ' ' {
		write--
	}
	return b[:write]
}
