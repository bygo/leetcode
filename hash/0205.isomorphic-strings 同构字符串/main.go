package main

// https://leetcode-cn.com/problems/isomorphic-strings

func isIsomorphic(s string, t string) bool {
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}
	for i := range s {
		if m1[s[i]] != 0 && m1[s[i]] != t[i] || m2[t[i]] != 0 && m2[t[i]] != s[i] {
			return false
		}
		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}
	return true
}
