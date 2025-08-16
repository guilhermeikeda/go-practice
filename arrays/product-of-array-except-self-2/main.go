package main

func main() {

}

func productExceptSelf(nums []int) []int {
	var answer []int = make([]int, len(nums))

	acc := 1
	for i := range nums {
		if i != 0 {
			acc = acc * nums[i-1]
		}

		answer[i] = acc
	}

	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 {
			suffix = suffix * nums[i+1] // calculate the suffix
		}

		answer[i] = suffix * answer[i] // result is the product of suffix and the prefix
	}

	return answer
}
