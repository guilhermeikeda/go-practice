package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfitEasy(prices)

	println(profit)
}

// Find the first lowest, then find the highest, than calculate the profit
func maxProfit(prices []int) int {
	buyPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		currentPrice := prices[i]
		if buyPrice > currentPrice {
			buyPrice = currentPrice
		}

		profit := prices[i] - buyPrice
		if profit > 0 {
			maxProfit += profit
			buyPrice = prices[i]
		}
	}

	return maxProfit
}

func maxProfitEasy(prices []int) int {
	buy := prices[0]
	var maxProfit int

	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		}

		profit := prices[i] - buy
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
