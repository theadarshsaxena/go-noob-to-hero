package main

// Given an integer array arr and an integer k, return the k most frequent elements. You may return the answer in any order.
// e.g. nums = [1,1,1,2,2,3], k = 2 then answer is [1,2]

func topKElements(arr []int, k int) []int {
	m := make(map[int]int)   // element -> frequency
	for _, v := range arr {
		m[v]++
	}
	freqarr := make([][]int, len(arr)+1)
	for i, v := range m {
		freqarr[v] = append(freqarr[v], i)
	}
	answer := []int{}
	kx := 0
	for i:=len(freqarr)-1; i>=0; i-- {
		if len(freqarr[i]) > 0 {
			answer = append(answer, freqarr[i]...)
			kx+=len(freqarr[i])
		}
		if kx>=k {
			break;
		}
	}
	return answer
}
// func main() {
// 	arr := []int{3,4,5,2,0,9,3,2,2}
// 	fmt.Println(topKElements(arr, 2))
// }