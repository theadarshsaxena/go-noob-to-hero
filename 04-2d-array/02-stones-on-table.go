package main

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