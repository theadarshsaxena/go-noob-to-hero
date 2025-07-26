package main

import (
	"slices"
)

func minDaysForBouquets(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	// binary search on the #days
	ans:=-1
	l,r := 1, slices.Max(bloomDay)
	for l<=r {
		mid := r - (r-l)/2
		if getDays(bloomDay, m, k, mid) == 1 {
			ans = mid
			r=mid-1
		} else {
			l = mid+1
		}
	}

	return ans
}

// This will return the possibility of creating bouquets
func getDays(arr []int, m, k, reqDays int) int {
	localLength := 0
	for _, v := range arr {
		if v<=reqDays {
			localLength++
		} else {
			localLength = 0
		}

		if localLength >= k {
			m--
			localLength=0
			if m==0 {
				return 1 // yes you can make bouquet in reqDays
			}
		}
	}
	return -1
}

// func main() {
// 	bloomDay := []int{7,7,7,7,12,7,7}
// 	m := 2
// 	k := 3
// 	fmt.Println(minDaysForBouquets(bloomDay, m, k))
// }