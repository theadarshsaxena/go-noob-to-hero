package main

func findKthPositive(arr []int, k int) int {
	n:=len(arr)
	if (arr[n-1]-len(arr)) < k {
		return (k+len(arr))
	}
    l,r:=0,n-1
	davedaar := -1
	mid := 0
	for l<=r {
		mid = l + (r-l)/2
		if arr[mid]-mid-1 > k {
			davedaar = mid
			r=mid-1
		} else {
			l=mid+1
		}
	}
	return davedaar+k
}
// func main() {
// 	arr := []int{5,6,7,8,9}
// 	k := 9
// 	fmt.Println(findKthPositive(arr, k))
// }

// TODO: Check the following approach, where the condition l=r is used for exiting the loop
/*
func findKthPositive(arr []int, k int) int {
    left, right := 0, len(arr)
    for left < right {
        mid := left + (right - left) / 2
        missing := arr[mid] - (mid + 1) // arr[mid] - (mid + 1) represents how many missing positives so far at index i
        if missing < k {
            left = mid + 1 // right part
        } else {
            right = mid // left part
        }
    }
    return left + k // returning number from indeces left
}
*/