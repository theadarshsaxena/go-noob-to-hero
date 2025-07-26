package main

/*
Problem Statement:
Given a non-negative integer n, find the integer square root of n using binary search. The integer square root is defined as the greatest integer m such that m*m <= n. If n is a perfect square, return its square root; otherwise, return -1.

Intuition Diagram:
Search Range: [l, r]
		   l         m         r
		   |---------|---------|
Each step:
- Calculate m = r - (r-l)/2
- Compare m*m with n
- Adjust l or r accordingly

This approach ensures O(log n) time complexity.
*/

func sqrtUsingBinarySearch(n int) int {
	if n==0 {
		return 0
	}
	l,r:=1,n
	for l<=r {
		m:=r-(r-l)/2
		if m*m==n {
			return m
		}
		if m*m > n {
			r = m-1
		} else {
			l = m+1
		}
	}
	return -1
}

// func main() {
// 	n:=1
// 	fmt.Println(sqrtUsingBinarySearch(n))
// }