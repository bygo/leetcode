package main

// https://leetcode-cn.com/problems/valid-anagram

func isAnagram(s string, t string) bool {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	for i := range t {
		m[t[i]]--
	}

	for _, cnt := range m {
		if cnt != 0 {
			return false
		}
	}
	return true
}
