package main

func searchInsertPosition(nums []int, target int) int {
	l,h,m:=0,len(nums)-1,(len(nums)-1)/2
	for l<h-1 {
		if (nums[m]==target) {
			return m
		}
		if nums[m] < target {
			l = m
		} else {
			h = m
		}
		m = (l+h)/2
	}
	if (nums[l]>=target) {
		return l
	} else if (nums[h]== target || target < nums[h] && target > nums[l]) {
		return h
	} else if (target > nums[h]) {
		return h+1
	}
	return -1
}

// func main() {
// 	nums := []int{1,3,5,6}
// 	fmt.Println(searchInsertPosition(nums, 5))
// }