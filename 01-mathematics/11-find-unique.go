package main

import "fmt"

// Given an array where every element appears exactly twice except one element that appears only once, find that single unique element.
func findUniqueIn(arr []int) int {
	num:=arr[0]
	for i:=1; i<len(arr); i++ {
		num ^= arr[i]
	}
	return num
}

func main() {
	arr := []int{2, 3, 5, 4, 5, 3, 4}
	fmt.Println(findUniqueIn(arr))
}