package main

import "fmt"

func bitwiseDemo(a, b int) {
	fmt.Printf("a = %d (%08b), b = %d (%08b)\n", a, a, b, b)

	fmt.Printf("a & b  = %d (%08b)\n", a&b, a&b)     // AND
	fmt.Printf("a | b  = %d (%08b)\n", a|b, a|b)     // OR
	fmt.Printf("a ^ b  = %d (%08b)\n", a^b, a^b)     // XOR
	fmt.Printf("a << 1 = %d (%08b)\n", a<<1, a<<1)   // Left shift (multiply by 2)
	fmt.Printf("b >> 1 = %d (%08b)\n", b>>1, b>>1)   // Right shift (divide by 2)
}

// a = 5 (00000101), b = 3 (00000011)
// a & b  = 1 (00000001)
// a | b  = 7 (00000111)
// a ^ b  = 6 (00000110)
// a << 1 = 10 (00001010)
// b >> 1 = 1 (00000001)