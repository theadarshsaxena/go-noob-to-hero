package main

func longestCommonPrefix(strs []string) string {
	leastLength := len(strs[0])
	for i := range strs {
		if leastLength > len(strs[i]) {
			leastLength = len(strs[i])
		}
	}
	result := []byte{}
	for i:=0; i<leastLength; i++ {
		x:= strs[0][i]
		for j:=1;j<len(strs); j++ {
			if strs[j][i] != x {
				return string(result)
			}
		}
		result = append(result, x)
	}
	return string(result)
}

// func main() {
// 	strs := []string{"flower","flow","flight"}
// 	fmt.Println(longestCommonPrefix(strs))
// }