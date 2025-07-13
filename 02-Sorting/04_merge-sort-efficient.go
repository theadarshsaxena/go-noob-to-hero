package main

func mergeSort(arr []int, l, h int, temp []int) {
    if (l==h) { // base case
        return
    }
    mid := (l+h)/2
    mergeSort(arr, l, mid, temp)
    mergeSort(arr, mid+1, h, temp)
    merge(arr, temp, l, mid, h)
}
func merge(arr, temp []int, l, m, h int) {
    p := 0 // Will point to temp
    arrOriginalStart := l
    end := h-l
    i, j := l, m+1
    for i<=m && j<=h {
        if (arr[i]<arr[j]) {
            temp[p] = arr[i]
            i++
        } else {
            temp[p] = arr[j]
            j++
        }
        p++
    }
    // assign rest of the elements
    if(j<=h) {
        for (j<=h) {
            temp[p] = arr[j]
            j++
            p++
        }
    } else {
        for (i<=m) {
            temp[p] = arr[i]
            i++
            p++
        }
    }
    // Assign back to original array
    for a:=0; a<=end; a++ {
        arr[arrOriginalStart] = temp[a]
        arrOriginalStart++
    }
}
// func main() {
//   arr := []int{4,5,2,0,1,2}
//   temp := make([]int, len(arr))
//   mergeSort(arr, 0, len(arr)-1, temp)
//   fmt.Println(arr)
// }