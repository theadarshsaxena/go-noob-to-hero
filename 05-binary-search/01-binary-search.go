package main

func binarySearch(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, (0+len(nums)-1)/2
	for mid <= high && mid >= low {
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid+1
		}
		mid = (low + high)/2
	}
	return -1
}

// func main() {
// 	arr := []int{-1,0,3,5,9,12}
// 	fmt.Println(binarySearch(arr, 9))
// }