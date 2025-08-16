package main

// * Rearrange Array Elements by Sign
// * There’s an array ‘A’ of size ‘N’ with an equal number of positive and negative elements.
// * Without altering the relative order of positive and negative elements, you must return an array of alternately positive and negative values.

/*
Intuition for O(1) Space Solution:
- The idea is to rearrange the array in-place so that positive and negative numbers alternate, starting with a positive.
- For each index, check if the sign matches the expected (even index: positive, odd index: negative).
- If not, find the next element with the required sign and shift elements to place it at the current index, maintaining relative order.
- This uses no extra space but involves shifting elements, leading to higher time complexity.

Intuition for O(n) Space Solution:
- Separate all positive and negative numbers into two new arrays, preserving their original order.
- Then, merge them back into the original array by alternating elements from the positive and negative arrays.
- This approach uses extra space but is faster since it avoids shifting elements.
*/

// Note: Start the array with positive elements.
// Note: Following commented code is having O(1) space complexity
// func arrangePosNeg(arr []int) {
// 	for i:=0; i<len(arr); i++{
// 		if i%2==0 {
// 			if arr[i] < 0 {
// 				p := findNextPositiveIndex(arr, i)
// 				shiftAndSwap(arr, i, p)
// 			}
// 		} else {
// 			if arr[i] >= 0 {
// 				n := findNextNegIndex(arr, i)
// 				shiftAndSwap(arr, i, n)
// 			}
// 		}
// 	}
// }

// func findNextPositiveIndex(arr []int, i int) int {
// 	i++
// 	for i < len(arr) {
// 		if arr[i] > 0 {
// 			return i
// 		}
// 		i++
// 	}
// 	return -1
// }

// func findNextNegIndex(arr []int, i int) int {
// 	i++
// 	for i < len(arr) {
// 		if arr[i] < 0 {
// 			return i
// 		}
// 		i++
// 	}
// 	return -1
// }

// func shiftAndSwap(arr []int, i, p int) {
// 	temp:=arr[p]
// 	for iter:=p; iter>i; iter-- {
// 		arr[iter] = arr[iter-1]
// 	}
// 	arr[i]=temp
// }

// Note: Following is having O(n) Space but better time complexity than above

func arrangePosNeg(arr []int) {
	var sizeOfPos int
	if len(arr)%2==0 {
		sizeOfPos = len(arr)/2
	} else {
		sizeOfPos = len(arr)/2+1
	}
	narr := make([]int, len(arr)/2)
	parr := make([]int, sizeOfPos)
	n,p:=0,0
	for i:=0; i<len(arr); i++ {
		if arr[i] < 0 {
			narr[n]= arr[i]
			n++
		} else {
			parr[p] = arr[i]
			p++
		}
	}
	// reassign to original array
	x:=0
	for x=0; x<n; x++ {
		arr[2*x] = parr[x]
		arr[2*x+1] = narr[x]
	}
	if len(arr) % 2!=0 {
		arr[2*x]=parr[x]
	}
}
	

// func main() {
// 	arr := []int{28,-41,22,-8,-37,46,35,-9,18,-6,19,-26,-37,-10,-9,15,14,31}
// 	arrangePosNeg(arr)
// 	fmt.Println(arr)
// }