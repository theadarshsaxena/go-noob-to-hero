package main

// https://leetcode.com/problems/remove-outermost-parentheses/description/

/* Intuition
First thing is we have two characters in the string, so, we can use a variable
which can act as a stack. ( means stack++ and ) means stack--

So, for each character, check if char is ( or ) and
if ( then check if it is outermost or not -> if outermost then don't add to result else add

and if ) then, check if stack == 1 -> means ) is the last outermost -> so don't add to result
*/

func removeOuterParentheses(s string) string {
	numOpened := 0
	result := []byte{}
	for _, v := range s {
		if v == rune('(') {
			if numOpened != 0 {
				result = append(result, '(')
			}
			numOpened++
		} else {
			if numOpened != 1 {
				result = append(result, ')')
			}
			numOpened--
		} 
	}
	resultString := string(result)
	return resultString
}

// func main() {
// 	str := "(()())(())"
// 	fmt.Println(removeOuterParentheses(str))
// }