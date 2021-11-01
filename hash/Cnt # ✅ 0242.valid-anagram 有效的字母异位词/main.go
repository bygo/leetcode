package main

// https://leetcode-cn.com/problems/valid-anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}

	for _, cnt := range m {
		if cnt != 0 {
			return false
		}
	}
	return true
}
