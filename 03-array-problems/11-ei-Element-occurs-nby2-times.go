package main

// Element that occurs more than N/2 times
// * ⚡ Boyer-Moore Voting Algorithm — THE Genius Intuition

// We don’t need to count everyone. Just the one that survives opposition.
// 🤖 Mental Model:
//     Keep a count and a candidate
//     Go through the array:
//         If count == 0, pick a new candidate
//         If current element == candidate → count++
//         Else → count--
//     🎯 The majority element survives all challenges, because it appears more than all the others combined

func majorityElement(arr []int) int {
	majorityElement, currentCount := 0, 0
	for _, v := range arr {
		if currentCount == 0 {
			majorityElement = v
			currentCount++
		} else {
			if v == majorityElement {
				currentCount++
			} else {
				currentCount--
			}
		}
	}
	return majorityElement
}

// func main() {
// 	arr := []int{2,3,2,2,1,1,2}
// 	fmt.Println(majorityElement(arr))
// }