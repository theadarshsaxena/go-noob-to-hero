package main

func isEven(num int) bool {
	// If the last bit (LSB) is 1, the number is odd, else even
	return (num&1)==0
}
// func main() {
// 	if isEven(23) {
// 		fmt.Println("Even")
// 	} else {
// 		fmt.Println("odd")
// 	}
// }