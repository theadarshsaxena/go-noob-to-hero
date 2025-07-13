package main

// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.


func commonElements(nums1 []int, nums2 []int) []int {
	// add all elements of nums1 to array
	s := make(map[int]bool)
	for _, v := range nums1 {
		s[v] = true
	}
	m := make(map[int]bool)
	for _, v := range nums2 {
		if s[v] {
			m[v] = true
		}
	}
	var resultArr []int
	for i := range m {
		resultArr = append(resultArr, i)
	}
	return resultArr
}

// func main() {
// 	arr := []int{1,2,2,1}
// 	brr := []int{2,2}
// 	fmt.Println(commonElements(arr, brr))
// }