package main

import (
	"sort"
)

func threesum(arr []int) [][]int {
	sort.Ints(arr)
	var result [][]int
	for i:=0; i<len(arr)-2; i++ {
		if i>0 && arr[i]==arr[i-1] {
			continue
		}
		j, k := i+1, len(arr)-1
		for j<k {
			if arr[i]+arr[j]+arr[k] == 0 {
				result = append(result, []int{arr[i],arr[j],arr[k]})
				j++
				k--
				for j<k && arr[j]==arr[j-1] {
					j++
				}
				for j<k && arr[k]==arr[k+1] {
					k--
				}
			} else if (arr[i]+arr[j]+arr[k]) < 0 {
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
	return result
}

// func main() {
// 	arr := []int{2,-3,0,-2,-5,-5,-4,1,2,-2,2,0,2,-4,5,5,-10}
// 	fmt.Println(threesum(arr))
// }