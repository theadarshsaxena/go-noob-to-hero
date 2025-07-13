package main

import (
	"strconv"
)

func decimalToBinary(num int) string {
	// Direct way to convert
	binary := strconv.FormatInt(int64(num), 2)
	return binary

	// TODO: Write logic by yourself to convert the decimal to binary
}
// func main() {
// 	fmt.Println(decimalToBinary(21))
// }