package main

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte)
	for i := range s {
		if v, ok := m[s[i]]; ok {
			if v == t[i]{
				continue
			} else {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}
	s, t = t, s
	m2 := make(map[byte]byte)
	for i := range s {
		if v, ok := m2[s[i]]; ok {
			if v == t[i]{
				continue
			} else {
				return false
			}
		} else {
			m2[s[i]] = t[i]
		}
	}
	return true
}

// func main() {
// 	fmt.Println(isIsomorphic("badc", "baba"))
// }

// TODO: try to incorporate the logic in single loop instead of two separate.