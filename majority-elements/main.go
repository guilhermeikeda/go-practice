package main

/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.

Ex: [6, 5, 5] -> 5

Solve in O(1)
*/
func majorityElements(nums []int) int {
	result := make(map[int]int, len(nums))
	for _, value := range nums {
		result[value]++
		if result[value] > len(nums)/2 {
			return value
		}
	}

	return 0
}
