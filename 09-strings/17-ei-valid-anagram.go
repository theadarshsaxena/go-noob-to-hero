package main

// Given two strings s and t, return true if t is an of s, and false otherwise.
// An anagram is a word or phrase formed by rearranging the letters of a
// different word or phrase, using all the original letters exactly once.

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)  // element -> frequency
	// add all elements of s
	for _, ch := range s {
		m[ch]++
	}
	// iterate through string t to see if element not exist in map
	for _, ch := range t {
		if freq, ok := m[ch]; ok && freq > 0 {
			m[ch]--
		} else {
			return false
		}
	}
	return true
}