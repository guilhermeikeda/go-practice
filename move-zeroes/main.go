package main

func main() {
	moveZeroes([]int{0, 0, 2, 0, 1})
}

func moveZeroes(nums []int) {
	lastNonZeroFoundAt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}
