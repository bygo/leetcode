package main

// https://leetcode-cn.com/problems/isomorphic-strings

func isIsomorphic(s string, t string) bool {
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}
	l := len(s)
	for i := 0; i < l; i++ {
		ch1 := s[i]
		ch2 := t[i]
		if m1[ch1] != 0 && m1[ch1] != ch2 || m2[ch1] != 0 && m2[ch2] != ch1 {
			return false
		}
		m1[ch1] = ch2
		m2[ch2] = ch1
	}

	return true
}
