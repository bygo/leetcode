package main

// https://leetcode-cn.com/problems/dKk3P7/

func isAnagram(s string, t string) bool {
	if s == t || len(s) != len(t) {
		return false
	}
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	for i := range t {
		m[t[i]]--
	}

	for b := range m {
		if m[b] != 0 {
			return false
		}
	}
	return true
}
