package main

import (
	"sort"
)

// https://leetcode.com/problems/3sum/description/

func threeSum(arr []int) [][]int {
	resultArray := make([][]int, 0)
	for i, num := range arr {
		requiredSum := 0-num
		m := make(map[int]int)
		for j:=i+1; j<len(arr); j++ {
			if index, exist := m[requiredSum-arr[j]]; exist {
				tempArray := []int{num, arr[j], arr[index]}
				sort.Ints(tempArray)  // Note: this is how sort the integer array
				if !contains(resultArray, tempArray) {
					resultArray = append(resultArray, tempArray)
				}
			}
			m[arr[j]]=j
		}
	}
	return resultArray
}

func contains(slice [][]int, element []int) bool {
	for _, arr := range slice {
		if len(arr) != len(element) {
			continue
		}
		match := true
		for i := range arr {
			if arr[i] != element[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// func main() {
// 	arr := []int{-1,0,1,2,-1,-4}
// 	fmt.Println(threeSum(arr))
// }