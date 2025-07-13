package main

import (
	"fmt"
	"math"
	"sort"
)

func printDivisors(num int) {
	resultArr := []int{}
	sqrt := int(math.Sqrt(float64(num)))
	for i:=2; i<=sqrt; i++ {
		if num%i == 0 {
			resultArr = append(resultArr, i)
			if i!=num/i {  // TODO: Check this condition, this is crucial to minimize the time
				resultArr = append(resultArr, num/i)
			}
		}
	}
	sort.Ints(resultArr)
	fmt.Println(resultArr)
}

// func main() {
// 	printDivisors(64)
// }