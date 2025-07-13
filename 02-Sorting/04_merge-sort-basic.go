package main

func mergeSortInefficient(arr []int) []int {
    // base case
    if len(arr)<=1 {
        return arr
    }
    mid := len(arr)/2

    firstHalf := mergeSortInefficient(arr[:mid])
    secondHalf := mergeSortInefficient(arr[mid:])
    return mergeInefficient(firstHalf, secondHalf)
}

func mergeInefficient(first []int, second []int) []int {
    xarr := make([]int, 0, len(first)+len(second))
    a := 0
    i, j := 0, 0
    for i<len(first) && j<len(second) {
        if (first[i]<second[j]) {
            xarr = append(xarr, first[i])
            i++
        } else {
            xarr = append(xarr, second[j])
            j++
        }
        a++
    }
    xarr = append(xarr, first[i:]...)
    xarr = append(xarr, second[j:]...)
    return xarr
}

// func main() {
//   arr := []int{4,5,2,0,1,2}
//   sortedarr := mergeSortInefficient(arr)
//   fmt.Println(sortedarr)
// }