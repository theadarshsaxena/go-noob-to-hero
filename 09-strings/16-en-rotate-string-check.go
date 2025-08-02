package main

func rotateString(s string, goal string) bool {
	n := len(goal)
	if len(s) != n {
		return false
	}
	xarr := []int{}
	for i := range goal {
		if s[0] == goal[i] {
			xarr = append(xarr, i)
		}
	}
	for _, x := range xarr {
		result := false
		for i := range s {
			if s[i] != goal[(i+x)%n] {
				result = false
				break
			} else {
				result = true
			}
		}
		if result {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(rotateString("abcde", "deabc"))
// }