package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
    splited := strings.Fields(s)
	slices.Reverse(splited)
	return strings.Join(splited, " ")
}

// func main() {
// 	str := "a good   example"
// 	fmt.Println(reverseWords(str))
// }