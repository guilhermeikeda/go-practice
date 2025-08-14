package main

func main() {

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// First pass: prefix products
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// Second pass: suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
