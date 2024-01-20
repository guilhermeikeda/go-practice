package main

import "fmt"

func main() {
	input := []int{1, 1, 1, 2, 2, 3}
	k := 5

	result := removeDuplicates(input)
	fmt.Printf("%d is equal to %d: %v", k, result, k == result)
}

func removeDuplicates(nums []int) int {
	var bookmark int = 1
	var gate bool = true

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[bookmark] = nums[i]
			bookmark++
			gate = true
		} else if gate {

			nums[bookmark] = nums[i]
			bookmark++
			// found a repeated element
			gate = false
		}
	}
	return bookmark
}
