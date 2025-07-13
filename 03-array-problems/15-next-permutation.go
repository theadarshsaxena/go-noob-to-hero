package main

import (
	"sort"
)

func nextPermutation(arr []int) {
	// find pivot
	// for i:=len(arr)-2; i>=0; i-- {
	// 	if arr[i] < arr[i+1]{
	// 		// ith element is pivot here
	// 		//find the smallest in arr[i+1:] greater than arr[i] and swap with arr[i]
	// 		smallest := math.MaxInt
	// 		smallestElementIndex := -1
	// 		var j int
	// 		for j=i+1; j<len(arr); j++ {
	// 			if arr[j]>arr[i] {
	// 				smallest = min(smallest, arr[j])
	// 				smallestElementIndex = j
	// 			}
	// 		}
	// 		arr[i], arr[smallestElementIndex] = arr[smallestElementIndex], arr[i]
	// 		// sort the array arr[i+1:] -> this is essentially a reversing of the array
	// 		sort.Ints(arr[i+1:])
	// 		break
	// 	}
	// }
	// if i<0 {
		
	// }
	if len(arr)==1 {
		return
	}
	pivot := 0
	for i:= len(arr)-1; i>0; i-- {
		if (arr[i] > arr[i-1]) {
			pivot = i-1
			break
		}
		if i==1 {
			sort.Ints(arr)
			return
		}
	}
	// find the just bigger element
	m := pivot+1
	for j:=pivot+1; j<len(arr); j++ {
		if arr[pivot] < arr[j] {
			if arr[j] < arr[m] {
				m = j
			}
		}
	}
	// Replace the just greater element with pivot
	arr[m], arr[pivot] = arr[pivot], arr[m]
	// Sort the element just after the pivot
	sort.Ints(arr[pivot+1:])
}
// func main() {
// 	tests := []struct {
// 		input    []int
// 		expected []int
// 	}{
// 		{[]int{1, 2, 3}, []int{1, 3, 2}},
// 		{[]int{3, 2, 1}, []int{1, 2, 3}},
// 		{[]int{1, 1, 5}, []int{1, 5, 1}},
// 		{[]int{1, 3, 2}, []int{2, 1, 3}},
// 		{[]int{2, 3, 1}, []int{3, 1, 2}},
// 	}

// 	for _, tt := range tests {
// 		arr := make([]int, len(tt.input))
// 		copy(arr, tt.input)
// 		fmt.Println(arr)
// 		nextPermutation(arr)
// 		fmt.Println(arr)
// 		if !reflect.DeepEqual(arr, tt.expected) {
// 			fmt.Printf("FAIL: nextPermutation(%v) = %v; want %v\n", tt.input, arr, tt.expected)
// 		} else {
// 			fmt.Printf("PASS: nextPermutation(%v) = %v\n", tt.input, arr)
// 		}
// 	}
// }