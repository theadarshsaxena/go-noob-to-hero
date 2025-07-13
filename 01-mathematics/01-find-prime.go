package main

import (
	"math"
)

// Prime nums are greater than 1 and only have 1 divisor other than 1
// Find a prime number
func findPrime(num int) bool {
	if num==0 || num == 1 {
		return false
	}
	for i:=2; i<int(math.Sqrt(float64(num))); i++ {
		if num%i==0 {
			return false
		}
	}
	return true
}

// Find prime numbers till n
// Using sieve of eratosthenes

func sieve(num int) int {
	baseArr := make([]int, num+1)
	result :=0
	for i := 2; i<num; i++ {
		if baseArr[i]==0{
			result++
		}
		for j:=i*2; j<=num; j=j+i {
			baseArr[j]=1
		}
	}
	return result
}

// TODO: Following approach is way faster, check this.

	// if n <= 2 {
    //     return 0
    // }

    // entries := make([]bool, n >> 1)
    // primeCount := len(entries)

    // for i := 1; i * i < (n >> 1); i++ {
    //     if !entries[i] {
    //         for j := 1 + (i << 1) + i; j < len(entries); j += 1 + (i << 1) {
    //             if !entries[j]{
    //                 entries[j] = true
    //                 primeCount--
    //             }
    //         }
    //     }
    // }

    // return primeCount

// func main() {
// 	var num = 21
// 	fmt.Println(findPrime(num))
// 	fmt.Printf("prime numbers till %d are %v",num,sieve(num))
// }