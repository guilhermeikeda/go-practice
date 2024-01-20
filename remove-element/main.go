package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	k := 5

	result := removeElement(nums, val)
	fmt.Printf("%d is equal to %d: %v", k, result, k == result)
}

// nums input, and we should remove all occurances of val
// return the K elements
// func removeElement(nums []int, val int) int {
// 	var k int

// 	// Scan the entire array
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == val {
// 			continue
// 		}

// 		nums[k] = nums[i]
// 		k++
// 	}

// 	return k
// }

func removeElement(nums []int, val int) int {

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	return j
}
