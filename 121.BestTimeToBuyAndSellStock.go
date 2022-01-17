package leetcode

func maxProfit(prices []int) int {
	var profit int
	var minPrice int = prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
