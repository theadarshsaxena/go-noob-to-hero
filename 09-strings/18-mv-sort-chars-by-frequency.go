package main

import (
	"sort"
	"strings"
)

/*
https://leetcode.com/problems/sort-characters-by-frequency/description/
Given a string s, sort it in decreasing order based on the frequency of the characters.
The frequency of a character is the number of times it appears in the string.

Return the sorted string. If there are multiple answers, return any of them.
E.g.
Input: s = "tree"
Output: "eert"
Explanation: 'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/

func frequencySort(s string) string {
    m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	type freqCount struct {
		character byte
		freq int
	}
	freqArray := []freqCount{}
	for i, v := range m {
		freqArray = append(freqArray, freqCount{i,v})
	}
	sort.Slice(freqArray, func(i, j int) bool {return freqArray[i].freq > freqArray[j].freq})
	var result strings.Builder
	for _, v := range freqArray {
		result.WriteString(strings.Repeat(string(v.character), v.freq))
		// result = append(result, strings.Repeat(v.character, v.freq))
	}
	return result.String()
}

// func main() {
// 	fmt.Println(frequencySort("tree"))
// }

// TODO: check the following logic, much faster & simple
/*
func frequencySort(s string) string {
    freq := make([]int, 255)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})

	// fmt.Printf("s: %s, ans: %s\n", s, string(b))
	return string(b)
}
*/