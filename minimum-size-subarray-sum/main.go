package main

import "fmt"

func main() {
	target := 11
	nums := []int{1, 1, 1, 1, 7}
	value := minSubArrayLen(target, nums)
	fmt.Println(value)
}

func minSubArrayLen(target int, nums []int) int {
	out := 0
	start, sum := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target { // found the max possibility, even summing all the values the sum is lower, than it's impossibel
			if out == 0 || (i-start+1) < out {
				out = i - start + 1
			}

			sum -= nums[start]
			start++
		}
	}
	return out
}
