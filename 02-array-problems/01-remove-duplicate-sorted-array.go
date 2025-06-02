package main

// RemoveDuplicatesSortedArray removes duplicates from a sorted array in-place
// and returns the array.
func RemoveDuplicatesSortedArray(arr []int) []int {
	i, j := 0, 0
	for j < len(arr) {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
		j++
	}
	// Remove extra elements from the array
	arr = arr[:i+1]
	return arr
}

// func main() {
// 	arr := []int{2, 2, 3, 3, 3, 5}
// 	arr = RemoveDuplicatesSortedArray(arr)
// 	fmt.Println(arr)
// }