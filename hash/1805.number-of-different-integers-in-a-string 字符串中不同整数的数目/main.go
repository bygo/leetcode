package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

func numDifferentIntegers(word string) int {
	m := map[string]struct{}{}
	word += "#"
	var cur []byte
	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			cur = append(cur, word[i])
		} else {
			l1 := len(cur)
			if 0 < l1 {
				i := 0
				for cur[i] == '0' && i < len(cur)-1 {
					i++
				}
				m[string(cur[i:])] = struct{}{}
				cur = []byte{}
			}
		}
	}
	return len(m)
}
