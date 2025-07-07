package main

// Given unsorted array with repeated elements, you need to in-place remove the duplicates from the array
// and the assign zero at the end of the array for the empty places created.
// Space complexity = O(1)
// Time complexity = O(n)

// e.g. 3,4,9,3,7,3,5 --> [3 4 9 7 5 0 0]

func removeDuplicates(arr []int) {
	m := make(map[int]bool)  // element -> index
	ci := 0
	for i := range arr {
		if _, ok := m[arr[i]]; ok {
			continue
		}
		if i!= ci {
			arr[ci] = arr[i]
		}
		m[arr[i]] = true
		ci++
	}

	for ci < len(arr) {
		arr[ci] = 0
		ci++
	}
}

// func main() {
// 	arr := []int{3,4,9,3,7,3,5}
// 	removeDuplicates(arr)
// 	fmt.Println(arr)
// }