package main

// https://leetcode-cn.com/problems/find-the-difference

func findTheDifference(s string, t string) byte {
	m := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		m[ch]++
	}

	for i := range t {
		ch := t[i] - 'a'
		if m[ch] == 0 {
			return ch + 'a'
		}
		m[ch]--
	}
	return 0
}
