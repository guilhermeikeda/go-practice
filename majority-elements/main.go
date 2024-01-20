package main

func majorityElements(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 1
		current := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if current == nums[j] {
				count++
			}
		}

		if count > len(nums)/2 {
			return current
		}
	}

	return 0
}
