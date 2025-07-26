package main

import "fmt"
func shipWithinDays(weights []int, days int) int {
    l, r:=0,0
	for i := range weights {
		if weights[i] > l {
			l = weights[i]
		}
		r += weights[i]
	}
	mid:=0
	ans := -1
	for l<=r {
		mid = l+(r-l)/2
		if getDaysToLoad(weights, mid) <= days {
			ans = mid
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return ans
}

func getDaysToLoad(weights []int, capacity int) int {
	sum := 0
	days := 0
	for i := range weights {
		if sum+weights[i] <= capacity {
			sum+=weights[i]
		} else {
			sum=weights[i]
			days++
		}
	}
	if sum!=0 {
		days++
	}
	return days
}

func main() {
	weights := []int{1,2,3,4,5,6,7,8,9,10}
	days := 5
	fmt.Println(shipWithinDays(weights, days))
}