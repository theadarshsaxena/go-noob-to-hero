// IN PROGRESS

package main

func check(nums []int) bool {
    minIndex := findMinInArray(nums)
	rotateByN(nums, minIndex)
	return checkSorted(nums)
}

func checkSorted(arr []int) bool {
	for i:=0; i<len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func findMinInArray(arr []int) int {
    i := 0
	for i<len(arr)-1{
		if arr[i] <= arr[i+1] {
			i++
		} else {
			i++
			break
		}
	}
	if i==len(arr)-1 {
		return 0
	}
	return i
}

func rotateByN(arr []int, N int) {
    j, elementCount := 0, 0
	for {
		temp:= arr[j]
		i := j
		for {
			elementCount++
			arr[i] = arr[(N + i) % len(arr)]
			if (j== (i+N)%len(arr)) {
				arr[i] = temp
				break
			}
			i = (i + N)%len(arr)
		}
		if elementCount >= len(arr) {
			break
		}
		j++
	}
}

// func main() {
// 	arr := []int{1,2, 3}
// 	sorted := check(arr)
// 	fmt.Println(arr)
// 	fmt.Println(sorted)
// }