package main

func rotateby90degrees(arr [][]int) {
	// transpose
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr[0]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	// Reverse rows
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0])/2; j++ {
			arr[i][j], arr[i][len(arr[0])-j-1] = arr[i][len(arr[0])-j-1], arr[i][j]
		}
	}
}

// func main() {
// 	arr := [][]int{{3,4,9},{3,7,3},{5,6,2}}
// 	rotateby90degrees(arr)
// 	fmt.Println(arr)
// }