package main

import (
	"sort"
)

func fourSum(arr []int, target int) [][]int {
	var result [][]int
	sort.Ints(arr)
	for a := 0; a<len(arr)-3; a++ {
		//skip duplicate
		if a>0 && arr[a]==arr[a-1] {
			continue
		}
		for i:=a+1; i<len(arr)-2; i++ {
			if i > a+1 && arr[i]==arr[i-1] {
				continue
			}
			j, k := i+1, len(arr)-1
			for j<k {
				if arr[a]+arr[i]+arr[j]+arr[k] == target {
					result = append(result, []int{arr[a],arr[i],arr[j],arr[k]})
					j++
					k--
					for j<k && arr[j]==arr[j-1] {
						j++
					}
					for j<k && arr[k]==arr[k+1] {
						k--
					}
				} else if (arr[a]+arr[i]+arr[j]+arr[k]) < target {
					// skip duplicates and increase k
					j++
					for j<k && arr[j]==arr[j-1]{
						j++
					}
				} else {
					// skip duplicates and increase k
					k--
					for j<k && arr[k]==arr[k+1]{
						k--
					}
				}
			}
		}
	}
	return result
}

// func main() {
// 	arr := []int{1,0,-1,0,-2,2}
// 	fmt.Println(fourSum(arr, 0))
// }