package main

/*
Problem: Count the number of subarrays with given XOR K

Statement:
Given an array of integers and an integer K, count the total number of subarrays whose XOR is equal to K.

Intuition:
- Use prefix XOR logic to avoid recalculating subarray XORs.
- If `prefixXOR[j] ^ prefixXOR[i-1] == K`, then subarray(i...j) has XOR = K.
- Rearranged: `prefixXOR[i-1] = prefixXOR[j] ^ K`
- So, for each prefixXOR, check how many times `prefixXOR ^ K` has occurred before.
- Store prefixXOR frequencies in a map while traversing the array.

Time Complexity: O(n)
Space Complexity: O(n)
*/

func countSubarrayWithXorK(arr []int, k int) int {
	currXOR := 0
	freqMap := make(map[int]int)
	result := 0

	for _, num := range arr {
		currXOR ^= num

		if currXOR == k {
			result++
		}

		if count, ok := freqMap[currXOR^k]; ok {
			result += count
		}

		freqMap[currXOR]++
	}

	return result
}

// func main() {
// 	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
// 	k := 1
// 	fmt.Println(countSubarrayWithXorK(arr, k))
// }