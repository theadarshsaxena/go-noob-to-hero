package main

import (
	"fmt"
	"slices"
)

func aggressiveCows(arr []int, k int) int {
	slices.Sort(arr)
	allocationPossible := func(space, k int) bool {
		i, x:=1,0
		for i<len(arr) {
			if arr[i]-arr[x]>=space {
				x = i
				i++
				k--
			} else {
				i++
			}
			if k==1 {
				return true
			}
		}
		return false
	}
	l,r := 1, slices.Max(arr)
	ans := -1
	for l<=r {
		mid := l + (r-l)/2
		if allocationPossible(mid, k) {
			ans = mid
			l = mid+1
		} else {
			r = mid-1
		}
	}
	return ans
}

func main() {
	arr := []int{0,3,4,7,10,9}
	fmt.Println(aggressiveCows(arr, 4))
}