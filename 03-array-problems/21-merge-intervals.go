package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/merge-intervals/description/

func mergeIntervals(arr [][]int) [][]int {
	var result [][]int
	var s,e int  // store the start and end of entity for result array
	sort.Slice(arr, func(i, j int) bool {return arr[i][0] < arr[j][0]})
	for i:=0; i<len(arr); i++ {
		if i == 0 {
			s = arr[i][0]
			e = arr[i][1]
			continue
		}
		if e >= arr[i][0] {
			e = max(arr[i][1], e)
		} else {
			result = append(result, []int{s,e})
			s = arr[i][0]
			e = arr[i][1]
		}
	}
	result = append(result, []int{s,e})
	return result
}

func main() {
	arr := [][]int{{1,4},{0,4},{8,10},{15,18}}
	fmt.Println(mergeIntervals(arr))
}