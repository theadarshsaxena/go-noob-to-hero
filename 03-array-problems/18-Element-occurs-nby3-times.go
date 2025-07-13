package main

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

func majorityElementnby3(arr []int) []int {
	// Using moore's voting algorithm
	var n1, n2, c1, c2 int
	for _, num := range arr {
		switch {
		case c1 > 0 && n1 == num:
			c1++
		case c2>0 && n2 == num:
			c2++
		case c1==0:
			c1, n1 = 1, num
		case c2==0:
			c2, n2 = 1, num
		default:
			c1--
			c2--
		}
	}
	// Till now we have got max two elements which have occurred in the array but we don't know if they exist more than ⌊ n/3 ⌋ times.
	// So, let's count them

	c1, c2 = 0, 0
	for _, num := range arr {
		if num==n1 {
			c1++
		} else if num==n2 {
			c2++
		}
	}

	// Check if they exist more than  ⌊ n/3 ⌋ times
	var resultArr []int
	if c1*3 > len(arr) {
		resultArr = append(resultArr, n1)
	}
	if c2*3 > len(arr) {
		resultArr = append(resultArr, n2)
	}
	return resultArr
}

// func main() {
// 	arr := []int{3,2,3}
// 	fmt.Println(majorityElementnby3(arr))
// }