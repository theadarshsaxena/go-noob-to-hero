package main

// Given an array and sum k, find two integer values in the array which sums up to the k

func twoSum(arr []int, sum int) []int {
	m := make(map[int]int)
	for i, _ := range arr {
		if ind, exists := m[sum-arr[i]]; exists {
			return []int{ind, i}
		}
		m[arr[i]] = i
	}
	return []int{}
}

// func main() {
// 	arr := []int{2,7,34,5,66,7}
// 	fmt.Println(twoSum(arr, 9))
// }