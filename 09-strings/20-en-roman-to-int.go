package main

func romanToInt(s string) int {
	result := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

    for i := range s {
		if i == 0 {
			result += m[s[i]]
		} else {
			if m[s[i-1]] < m[s[i]] {
				result += m[s[i]] - 2*m[s[i-1]]
			} else {
				result += m[s[i]]
			}
		}
	}
	return result
}

// func main() {
// 	fmt.Println(romanToInt("MCMXCIV"))
// }


// TODO: Check the following faster implementation (minimize map call)
/*
var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	result := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		current := romanToArabic[string(s[i])]
		if current < prev {
			result -= current
		} else {
			result += current
		}
		prev = current
	}

	return result
}
*/