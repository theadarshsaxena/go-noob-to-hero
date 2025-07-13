package main

// Given an array, print all the elements which are leaders. A Leader is an element that is greater than all of the elements on its right side in the array.
// e.g.
//  arr = [4, 7, 1, 0]
// answer = 7 1 0
func leadersInArray(arr []int) []int {
	// intuition: iterate from right and keep a counter of max till that point and check if the current element greater than max or no
	max := -1
	resultArray := make([]int, 0)
	for i := len(arr)-1; i>=0; i-- {
		if arr[i] > max {
			resultArray = append(resultArray, arr[i])
			max = arr[i]
		}
	}
	n := len(resultArray)
	for i:=0; i<len(resultArray)/2; i++ {
		resultArray[i], resultArray[n-i-1] = resultArray[n-i-1], resultArray[i]
	}
	return resultArray
}

// func main() {
// 	arr := []int{4,7,1,0}
// 	fmt.Println(leadersInArray(arr))
// }