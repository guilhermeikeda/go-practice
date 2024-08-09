package main

import "fmt"

func main() {
	value := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Println(value)
}

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	window := 0
	l := len(nums)
	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}

	window = maxSum
	start := k

	for start < l {
		t := nums[start-k]
		window += nums[start] - t
		if window > maxSum {
			maxSum = window
		}
		start++
	}

	return float64(maxSum) / float64(k)
}
