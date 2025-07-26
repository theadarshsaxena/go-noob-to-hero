package main

import "fmt"

func upperBound(arr []int, t int) int {
	l,r:=0, len(arr)-1
	ans := -1
	for l<=r {
		m := l + (r-l)/2
		if arr [m] > t {
			r = m-1
		} else {
			ans = m
			l = m+1
		}
	}
	return ans
}

func main() {
	arr := []int{3,4, 5, 8, 9}
	k := 4
	fmt.Println(upperBound(arr, k))
}