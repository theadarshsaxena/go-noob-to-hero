package main

func leftRotateByOne(arr []int) {
	temp := arr[0]
	for i:=0; i<len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
}

// func main() {
// 	arr := []int{4,2,5,2,8,1}
// 	leftRotateByOne(arr)
// 	fmt.Println(arr)
// }