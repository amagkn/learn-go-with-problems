package max_profit

/*
MaxProfit

You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	currentMax := 0
	minPrice := prices[0]

	for _, n := range prices[1:] {
		if n < minPrice {
			minPrice = n
		} else {
			profit := n - minPrice

			if profit > currentMax {
				currentMax = profit
			}
		}
	}

	return currentMax
}
