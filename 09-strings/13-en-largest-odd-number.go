package main

import (
	"strconv"
)

func largestOddNumber(num string) string {
	oddFoundAt := -1
    for i, v := range num {
		if n,_:=strconv.Atoi(string(v)); n % 2 != 0 {
			oddFoundAt = i
		}
	}
	if oddFoundAt != -1 {
		return num[:oddFoundAt+1]
	} else {
		return ""
	}
}

// func main() {
// 	fmt.Println(largestOddNumber("35427"))
// }

// TODO: check the following better approach
/*
func largestOddNumber(num string) string {
    n := len(num)
    i := n - 1
    for ; i >= 0; i -= 1 {
        if (num[i] - '0') & 1 == 1 {
            break
        }
    }
    return num[:i + 1]
}
	*/