package main

// Given an array of int, find the longest consecutive sequence.
// e.g. for 100 2 100 1 3 4 -> 1 2 3 4 is the longest consecutive sequence.

func longestConsecutiveSequence(arr []int) int {
	m := make(map[int]bool)
	max := 0
	// add all to a set
	for i := range arr {
		m[arr[i]] = true
	}
	for i := range m {  // NOTE: pattern loop directly on the map
		tempCounter := 1
		if !m[i-1] {  // Note this pattern, how to directly use the condition out of map, if element not exist in map
			// Means this is the first element of the consecutive sequence.
			counter := 1
			for m[i+counter] {  // NOTE: again condition on "if element exist on map or not"
				tempCounter++
				counter++
			}
		}
		if max < tempCounter {
			max = tempCounter
		}
	}
	return max
}

// func main() {
// 	arr := []int{100, 2, 200, 1, 3, 4}
// 	fmt.Println(longestConsecutiveSequence(arr))
// }