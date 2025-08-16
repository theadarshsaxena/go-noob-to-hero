package main

func rightRotateByN(nums []int, n int) {
	counter := 0
	numsLen := len(nums)
	positiveMod := func (a, b int) int {
		return (a % b + b) % b
	}
	if (numsLen == n) {
		return
	}
	for i:=0; i < numsLen; i++ {
		temp := nums[i]
		j := i
		for {
			x := positiveMod(j-n, numsLen)
			nums[j] = nums[x]
			counter++
			if (positiveMod(j-n, numsLen)==i) {
				nums[j] = temp
				break
			}
			j=positiveMod(j-n, numsLen)
		}
		if counter >= numsLen {
			break
		}
	}
}

// func main() {
// 	arr := []int{1,2,3,4,5,6,7}
// 	rightRotateByN(arr, 1)
// 	fmt.Println(arr)
// }