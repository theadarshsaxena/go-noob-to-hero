package main

// given two string, find if they are anagram of not (anagram is the string we get by rearranging the letters)
func validAnagram(s, t string) bool {
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

// func main() {
// 	fmt.Println(validAnagram("anagram", "nagarbm"))
// }
