package main

func moveZerosToEnd(arr []int) {
	i, j := 0, 0
	for j<len(arr) {
		if (arr[j]==0) {
			j++
		} else {
			if (i!=j) {
				arr[i] = arr[j]
				i++
				j++
			} else {
				i++
				j++
			}
		}
	}
	for i<len(arr) {
		arr[i] = 0
		i++
	}
}

// func main() {
// 	arr := []int{2, 0, 2, 0, 0, 0, 3, 4, 1, 2, 0, 0}
// 	moveZerosToEnd(arr)
// 	fmt.Println(arr)
// }