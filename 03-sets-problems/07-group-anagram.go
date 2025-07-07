package main

import (
	"fmt"
)

// func groupAnagram(arr []string) [][]string {
// 	m := make(map[string][]int)
// 	for i := range arr {
// 		sbytes := []byte(arr[i])
// 		sort.Slice(sbytes, func(a, b int) bool { return sbytes[a] < sbytes[b] })
// 		k := string(sbytes)
// 		m[k] = append(m[k], i)
// 	}
// 	var s [][]string
// 	for _, v := range m {
// 		tempArray := make([]string, len(v))
// 		for j, k := range v {
// 			tempArray[j] = arr[k]
// 		}
// 		s = append(s, tempArray)
// 	}
// 	return s
// }



func groupAnagram(arr []string) [][]string {
	m := make(map[string][]string)
	for _, str := range arr {
		count := make([]int, 26)
		for _, ch := range str {
			count[ch-'a']++
		}
		// Create a unique key based on freq count
		key := fmt.Sprint(count)
		m[key] = append(m[key], str)
	}
	var s [][]string
	for _, v := range m {s = append(s, v)}
	return s
}


func main() {
	arr := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagram(arr))
}