package main

/*
Problem Statement:
Given a row of stones, each colored either red, green, or blue, represented as a string `s` of length `n` (where each character is 'R', 'G', or 'B'), you need to find the minimum number of stones to remove so that no two adjacent stones have the same color.

Intuition:
Iterate through the string and compare each stone with its previous one. If two adjacent stones have the same color, increment a counter for stones to remove. The final count gives the minimum number of stones to remove to ensure all adjacent stones are of different colors.
*/

func stoneOnTable(n int, s string) int {
	if n==1 {
		return 0
	}
	i := 0
	elementsToRemove := 0
	for j := range s {
		if j==0 {
			continue
		}
		if s[i] == s[j] {
			elementsToRemove++
			continue
		} else {
			i = j
		}
	}
	return elementsToRemove
}
// func main() {
// 	var a int
// 	var s string
// 	fmt.Scan(&a, &s)
// 	fmt.Println(stoneOnTable(a, s))
// }

// TO learn: iteration on string characters