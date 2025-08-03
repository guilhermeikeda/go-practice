package main

import "fmt"

func main() {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(result)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result = make([]bool, len(candies))
	for i, value := range candies {
		sum := value + extraCandies

		var isGreater = true
		for j, value2 := range candies {
			if i == j {
				continue
			}

			if sum < value2 {
				isGreater = false
				break
			}
		}

		result[i] = isGreater
	}

	return result
}
