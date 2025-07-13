package main

import "fmt"

func pascalTriangle(numRows int) [][]int {
	if numRows==0 {
		return [][]int{}
	}
	resultArr := [][]int{{1}}
	for i:=1; i<numRows; i++ {
		tempArray := make([]int, i+1)
		tempArray[0], tempArray[i] = resultArr[i-1][0], resultArr[i-1][i-1]
		for j := 1; j < i; j++ {
			tempArray[j] = resultArr[i-1][j-1] + resultArr[i-1][j]
		}
		resultArr = append(resultArr, tempArray)
	}
	return resultArr
}

func main() {
	fmt.Println(pascalTriangle(5))
}