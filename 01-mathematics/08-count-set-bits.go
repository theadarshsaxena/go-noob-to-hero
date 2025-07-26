package main

// logic:
// 	If you subtract 1 from a num, least significant set bit becomes 0 and all the bits following it becomes 1
//  e.g. 0100011100 - 1 => 0100011011
//  And do a & between (num-1) and (num) removes the least significant set bit.

func countSetBits(num int) int {
	counter := 0
	for num!=0 {
		num &= (num-1)
		counter++
	}
	return counter
}

// func main() {
// 	fmt.Println(countSetBits(29))
// }