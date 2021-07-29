package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string/

func numDifferentIntegers(word string) int {
	var cur int
	var m = map[int]struct{}{}
	for _, v := range word {
		if '0' <= v && v <= '9' {
			cur = cur*10 + int(v-48)
		} else {
			if cur != 0 {
				m[cur] = struct{}{}
			}
			cur = 0
		}
	}

	if cur != 0 {
		m[cur] = struct{}{}
	}
	return len(m)
}
