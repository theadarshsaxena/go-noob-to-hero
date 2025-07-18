package main

import "fmt"

//  Given a sorted array of N integers and an integer x, write a program to find the lower bound of x.
// Example 1:
// Input Format: N = 4, arr[] = {1,2,2,3}, x = 2
// Result: 1
// Explanation: Index 1 is the smallest index such that arr[1] >= x.

// Example 2:
// Input Format: N = 5, arr[] = {3,5,8,15,19}, x = 9
// Result: 3
// Explanation: Index 3 is the smallest index such that arr[3] >= x.

func lowerBound(arr []int, target int) int {
	low, high, mid := 0, len(arr)-1, (0+len(arr)-1)/2
	answer := -1
	for low <= high {
		if target <= arr[mid] {
			answer = mid
			high = mid -1
		} else {
			low = mid + 1
		}
		mid = (low+high)/2
	}
	return answer
}

func main() {
	arr := []int{3,5,8,15,19}
	fmt.Println(lowerBound(arr, 9))
}