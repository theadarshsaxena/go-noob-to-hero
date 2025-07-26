package main

//     Given a number n, return the XOR of all numbers from 0 to n.
//     Return the result in O(1) time.

//  Input: n = 6
//  Output: 0 ^ 1 ^ 2 ^ 3 ^ 4 ^ 5 ^ 6 = 7

func xorTillN(num int) int {
	// Finding the pattern, we can do it in O(1)
	modulo := num%4
	switch modulo {
	case 0: return num
	case 1: return 1
	case 2: return num+1
	case 3: return 0
	}
	return 0
}

// func main() {
// 	fmt.Println(xorTillN(7))
// }