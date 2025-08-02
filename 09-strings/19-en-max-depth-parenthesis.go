package main

/*
https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
Given a valid parentheses string s, return the nesting depth of s.
The nesting depth is the maximum number of nested parentheses.

Example 1:
Input: s = "(1+(2*3)+((8)/4))+1"
Output: 3
*/

func maxDepth(s string) int {
    depthCount := 0
	maxDepth := 0
	for _, v := range s {
		if v == '(' {
			depthCount++
			if maxDepth < depthCount {
				maxDepth = depthCount
			}
		} else if v == ')' {
			depthCount--
		}
	}
	return maxDepth
}

// func main() {
// 	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
// }