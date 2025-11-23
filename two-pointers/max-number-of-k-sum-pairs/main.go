package main

import "slices"

func main() {
}

func maxOperations(nums []int, k int) int {
	ptrA := 0
	ptrB := len(nums) - 1 // last index
	result := 0

	nums = sortAsc(nums)
	for ptrA < ptrB {
		sum := nums[ptrA] + nums[ptrB]
		if sum == k {
			ptrA++
			ptrB--
			result++
		} else if sum > k {
			ptrB--
		} else {
			ptrA++
		}
	}

	return result
}

func sortAsc(nums []int) []int {
	slices.Sort(nums)

	return nums
}
