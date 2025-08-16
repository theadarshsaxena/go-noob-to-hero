package main

// Set all the elements of a particular row and column where a 0 appears

func toxicZeroes(matrix [][]int) {
	// Intuition: Create two array to track which rows and columns need to set zero
	h,v:=make([]int, len(matrix)), make([]int, len(matrix[0]))
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				h[i] = 1
				v[j] = 1
			}
		}
	}
	// set the horizontal zeros
	for i:=0; i<len(h); i++ {
		if h[i]==1 {
			for j:=0; j<len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i:=0; i<len(v); i++ {
		if v[i]==1 {
			for j:=0; j<len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}

// func main() {
// 	// define the matrix
// 	matrix := [][]int{
// 		{2,3,4},
// 		{0,6,2},
// 		{2,3,0},
// 	}
// 	toxicZeroes(matrix)
// 	fmt.Println(matrix)
// }