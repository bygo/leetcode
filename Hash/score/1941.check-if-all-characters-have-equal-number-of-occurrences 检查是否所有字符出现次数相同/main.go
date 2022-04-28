package main

// https://leetcode-cn.com/problems/check-if-all-characters-have-equal-number-of-occurrences

func areOccurrencesEqual(s string) bool {
	m := map[byte]int{}
	base := s[0]
	for i := range s {
		m[s[i]] ++
	}
	for k := range m {
		if m[k] != m[base] {
			return false
		}
	}
	return true
}
