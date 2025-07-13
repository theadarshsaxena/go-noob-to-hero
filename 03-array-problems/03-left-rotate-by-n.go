package main

func leftRotateByN(arr []int, n int) {
	counter := 0
	arrLen := len(arr)
	if (arrLen == n) {
		return
	}
	for i:=0; i < arrLen; i++ {
		temp := arr[i]
		j := i
		for {
			arr[j] = arr[(j+n)%(arrLen)]
			counter++
			if ((j+n)%arrLen==i) {
				arr[j] = temp
				break
			}
			j=(j+n)%arrLen
		}
		if counter >= arrLen {
			break
		}
	}
}

// func main() {
// 	arr := []int{1,2,3,4,5,6,7}
// 	leftRotateByN(arr, 2)
// 	fmt.Println(arr)
// }