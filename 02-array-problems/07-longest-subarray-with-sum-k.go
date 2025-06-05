package main

import "fmt"

func getLongestSubarrayWithSumK(arr []int, k int) []int {
	l, e1, e2, i, j := 0, 0, 0, 0, 0
	for j<len(arr) {
		if getSum(arr, i, j) == k {
			if l < (j-i+1) {
				l = j-i+1
				e1 = i
				e2 = j
			}
			i++
			j++
		} else if getSum(arr, i, j) < k {
			j++
		} else {
			i++
		}
	}
	resultArr := make([]int, 0, e2-e1+1)
	resultArr = append(resultArr, arr[e1:e2+1]...)
	return resultArr
}

func getSum(arr[] int, i, j int) int {
	sum:=0
	for i<=j{
		sum=sum+arr[i]
		i++
	}
	return sum
}

func main() {
	arr := []int{2,3,5}
	fmt.Println(getLongestSubarrayWithSumK(arr, 5))
}