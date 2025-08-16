package main

// * You are given an array of prices where prices[i] is the price of a given stock on an ith day.
//
//  You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//  Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
//  Intuition:
//  Track the minimum price seen so far as you iterate through the array.
//  For each price, calculate the profit if you sold at that price (current price - min price so far).
//  Update the maximum profit if this profit is higher than previous maximum.
//  This way, you only need one pass through the array and constant space.

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