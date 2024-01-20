package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/
func twoSum(nums []int, target int) []int {
	result := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		requiredNumber := target - nums[i]
		if _, ok := result[requiredNumber]; ok { // In case the element is already in the map, we found our match
			return []int{result[requiredNumber], i}
		}

		result[nums[i]] = i // Otherwise, we store the visited number and its index
	}

	return []int{}
}
