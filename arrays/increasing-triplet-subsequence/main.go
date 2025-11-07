package main

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	increasingTriplet(nums)
}

func increasingTriplet(nums []int) bool {
	var first *int
	var second *int

	for _, num := range nums {
		if first == nil || num <= *first {
			first = &num
		} else if second == nil || num <= *second {
			second = &num
		} else {
			return true
		}
	}

	return false
}
