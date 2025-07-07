package main

import "fmt"

// This also works for 2d matrix which has uneven number of elements in rows
func rowWisePrint(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d", arr[i][j])
		}
		fmt.Println()
	}
}

// Diagonal traversal
func diagonalTraversal(arr [][]int) {
	for i:= 0; i<len(arr[0]); i++ {
		fmt.Printf("%d",arr[i][i])
	} 
}

// secondary diagonal traversal
func secondaryDiagonalTraversal(arr [][]int) {
	for i, j := 0, len(arr)-1; i<len(arr); i, j = i+1,j-1 {  // NOTE: i++, j-- isn’t valid in Go for multiple post-increments
		fmt.Printf("%d", arr[i][j])
	}
}

func TransposePrint(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d", arr[j][i])
		}
		fmt.Println()
	}
}


// func main() {
// 	arr := [][]int{{3,4,9},{3,7,3},{5,6,2}}
// 	rowWisePrint(arr)
// 	fmt.Println()
// 	diagonalTraversal(arr)
// 	fmt.Println()
// 	secondaryDiagonalTraversal(arr)
// 	fmt.Println()
// 	fmt.Println("Transpose")
// 	TransposePrint(arr)
// }