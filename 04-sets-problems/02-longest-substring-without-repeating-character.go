package main

// Given string, find the longest substring without repeating characters.
// e.g. abcba -> abc is the longest substring without repeating characters.

// func longestSubstringWithoutRepeatingChar(s string) int {
// 	m := make(map[rune]int)  // element -> index
// 	var (
// 		tempMax int
// 		max int
// 		lastTrackedIndex int
// 	)
// 	for i, ch := range s {
// 		if index, exists := m[ch]; exists {
// 			// Logic to remove the elements in the map till the last appearance of same character as arr[i]
// 			for j:=lastTrackedIndex; j<=index; j++ {
// 				delete(m, []rune(s)[j])
// 			}
// 			lastTrackedIndex = index+1
// 			tempMax = i-index-1
// 		}
// 		m[ch] = i
// 		tempMax++
// 		if tempMax > max {
// 			max = tempMax
// 		}
// 	}
// 	return max
// }

func longestSubstringWithoutRepeatingChar(s string) int {
	m := make(map[rune]int)  // element -> index
	var (
		max int // tracking the maximum subarray
		lt int // last tracked index
	)
	for i, ch := range s {
		if index, exists := m[ch]; exists {
			if lt <= index {
				lt = index + 1
			}
		}
		m[ch] = i
		if max < i - lt + 1 {
			max = i - lt + 1
		}
	}
	return max
}
// func main() {
// 	fmt.Println(longestSubstringWithoutRepeatingChar("abcabc"))
// }