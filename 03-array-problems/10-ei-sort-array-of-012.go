package main

// Dutch national flag algorithm
// If arr[mid] == 0 → swap arr[low] and arr[mid], then move both pointers forward
// If arr[mid] == 1 → just move mid forward
// If arr[mid] == 2 → swap arr[mid] and arr[high], then move high backward (don’t move mid!)
func sort012(arr []int) {
	l, m, h := 0, 0, len(arr)-1
	for m <= h {
		switch arr[m] {
		case 0:
			arr[m], arr[l] = arr[l], arr[m]
			m++
			l++
		case 1:
			m++
		case 2:
			arr[m], arr[h] = arr[h], arr[m]
			h--
		}
	}
}

// func main() {
// 	arr := []int{1,1,0,2,0,2,1,1,0,2,1}
// 	sort012(arr)
// 	fmt.Println(arr)
// }