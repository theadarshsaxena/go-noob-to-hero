package main

import (
	"slices"
)

func smallestDivisor(nums []int, threshold int) int {
    maxElement := slices.Max(nums)
	if threshold == len(nums) {
		return maxElement
	}
	l,r:=1,maxElement
	mid := 0
	ans := -1
	for l<=r {
		mid = l+(r-l)/2
		if getSumOfDivisions(nums, mid) <= threshold {
			ans = mid
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return ans
}

func getSumOfDivisions(nums []int, divisor int) int {
	sum := 0
	for i := range nums {
		// NOTE: this below line helps up avoid using int(math.Ceil(float64(num[i])/ float64(divisor)))
		// Here, we need 2/3 results in 1 and not zero.
		sum += (nums[i]+divisor-1)/divisor
	}
	return sum
}

// func main() {
// 	fmt.Println(smallestDivisor([]int{1,2,5,9}, 6))
// }