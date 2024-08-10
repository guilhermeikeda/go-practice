package main

import "fmt"

func main() {

	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var result [][]int = make([][]int, 2)

	var nums1Map map[int]int = make(map[int]int)
	var nums2Map map[int]int = make(map[int]int)

	for _, value := range nums1 {
		nums1Map[value] = value
	}

	for _, value := range nums2 {
		nums2Map[value] = value
	}

	for _, value := range nums1Map {
		if _, ok := nums2Map[value]; !ok {
			result[0] = append(result[0], value)
		}
	}

	for _, value := range nums2Map {
		if _, ok := nums1Map[value]; !ok {
			result[1] = append(result[1], value)
		}
	}

	return result
}
