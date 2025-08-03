package nextpermutation

func main() {

}

// Find the next lexicographically greater permutation
// If don't exist, return the lowest possible
func nextPermutation(input []int) []int {

	pivot := findPivot(input)
	if pivot == nil {
		return reverse(input)
	}

	successor := findSuccessor(*pivot, input)

	if successor == nil {
		return nil
	}

	pivotValue := input[*pivot]
	input[*pivot] = input[*successor]
	input[*successor] = pivotValue

	rest := reverse(input[*pivot+1:])
	beging := input[0 : *pivot+1]
	return append(beging, rest...)
}

// The pivot is the rightmost number that is lower than its following number
func findPivot(input []int) *int {
	for i := len(input) - 1; i > 0; i-- {
		if i-1 < 0 {
			return nil
		}

		if input[i-1] < input[i] {
			pivot := i - 1
			return &pivot
		}
	}

	return nil
}

// The successor is the rightmost number that is greater than the pivot
func findSuccessor(pivot int, input []int) *int {
	count := len(input) - 1

	for i := count; i >= 0; i-- {
		if input[i] == pivot {
			return nil
		}

		if input[i] > pivot {
			return &i
		}
	}
	return nil
}

func reverse(input []int) []int {
	i := 0
	j := len(input) - 1

	for i < j {
		aux := input[i]
		input[i] = input[j]
		input[j] = aux

		i++
		j--
	}

	return input
}

/*
Ex: [2, 4, 1, 7, 5, 0]
Want: []

Ex2:
[1, 2, 3, 4, 5] arr[i] > arr[i-1]?
[1, 2, 3, 5, 4]
[1, 2, 4, 3, 5]
[1, 2, 4, 3, 5]



*/
