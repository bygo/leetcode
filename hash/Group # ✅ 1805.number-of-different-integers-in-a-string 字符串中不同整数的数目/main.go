package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

func numDifferentIntegers(word string) int {
	m := map[string]struct{}{}
	var cur []byte
	f := func() {
		l1 := len(cur)
		if 0 < l1 {
			j := 0
			for cur[j] == '0' && j < l1-1 {
				j++
			}
			m[string(cur[j:])] = struct{}{}
			cur = []byte{}
		}
	}
	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			cur = append(cur, word[i])
		} else {
			f()
		}
	}
	f()
	return len(m)
}
