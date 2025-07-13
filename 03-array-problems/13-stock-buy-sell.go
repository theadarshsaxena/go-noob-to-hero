package main

//  You are given an array of prices where prices[i] is the price of a given stock on an ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func stockBuySell(arr []int) int {
	buy := int(^uint(0) >> 1) // assigns the maximum int value
	var sell, maxProfit int
	for i:=0; i<len(arr); i++ {
		if (arr[i] < buy) {
			buy = arr[i]
		}
		sell = arr[i]
		maxProfit = max(maxProfit, sell-buy)
	}
	return maxProfit
}

// func main() {
// 	arr := []int{7,1,5,3,6,2}
// 	fmt.Println(stockBuySell(arr))
// }