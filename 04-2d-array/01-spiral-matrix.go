package main

func sprial_matrix(arr [][]int) []int {
	//define directions
	dir := 0 // 0 -> right, 1->down, 2->left, 3->right
	count, i, j := 0, 0, 0
	// limiting varibles, that is, upto what point our pointer will go
	rl, dl, ll, ul := len(arr[0])-1, len(arr)-1, 0, 1
	totalElements := len(arr)*len(arr[0])
	resultArr := make([]int, totalElements)
	for count < totalElements {
		resultArr[count] = arr[i][j]
		count++

		// Go to next element,
		// if exist on same line, go there
		// else change direction and check if exist, go there.
		switch dir {
		case 0:
			if j<rl {
				j++
			} else {
				dir = (dir+1)%4
				rl=j-1
				i++
			}
		case 1:
			if i<dl {
				i++
			} else {
				dir = (dir+1)%4
				dl=i-1
				j--
			}
		case 2:
			if j>ll {
				j--
			} else {
				dir = (dir+1)%4
				ll=j+1
				i--
			}
		case 3:
			if i>ul {
				i--
			} else {
				dir = (dir+1)%4
				ul=i+1
				j++
			}
		}
	}
	return resultArr
}

// func main() {
// 	arr := [][]int{{3,4,9},{3,7,3},{5,6,2}}
// 	fmt.Println(sprial_matrix(arr))
// }