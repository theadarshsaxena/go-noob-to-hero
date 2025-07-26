package main

// Learn here:
//		 n | (1 << i) → Set i-th bit
//		 n & ^(1 << i) → Unset i-th bit
//		 n ^ (1 << i) → Toggle i-th bit

// set the bit at ith position in given number
func setbit(num int, i int) int {
	return num | (1<<i)
}

// unset the bit at ith position in given number
func unsetBit(num int, i int) int {
	return num & ^(1<<i)
}

// toggle bit at ith position
func toggleBit(num int, i int) int {
	return num ^ (1<<i)
}
// func main() {
// 	n := 10       // 1010
// 	pos := 2      // working on 2nd bit from right (0-indexed)

// 	fmt.Printf("Original: %b\n", n)
// 	fmt.Printf("Set bit %d: %b\n", pos, setbit(n, pos))
// 	fmt.Printf("Unset bit %d: %b\n", pos, unsetBit(n, pos))
// 	fmt.Printf("Toggle bit %d: %b\n", pos, toggleBit(n, pos))
// }