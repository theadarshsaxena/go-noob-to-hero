package main

// * find the maximum subarray sum in array

// * ⚡ Kadane’s Intuition — The Golden Insight

// Here’s the mental model you need to train into your brain:
//     “If a subarray is hurting me (i.e. total sum becomes negative), I should ditch it and start fresh from the next element.”
// Basically:
//     At every point, decide:
//         “Do I extend my current subarray?”
//         Or “Do I start a new one from this index?”

func maxSum(arr []int) int {
	currentSum, maxSum := arr[0], arr[0]
	for _, v := range arr {
		if currentSum < 0 {
			currentSum = v
		} else {
			currentSum += v
		}
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

// func main() {
// 	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
// 	fmt.Println(maxSum(arr))
// }