package main

// https://leetcode-cn.com/problems/check-if-all-characters-have-equal-number-of-occurrences

func areOccurrencesEqual(s string) bool {
	m := map[byte]int{}
	for i := range s {
		m[s[i]] ++
	}

	first := 0
	for i := range m {
		if first == 0 {
			first = m[i]
		} else if first != m[i] {
			return false
		}
	}
	return true
}
