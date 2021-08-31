package main

// https://leetcode-cn.com/problems/longest-palindrome

func longestPalindrome(s string) int {
	m := map[byte]int{}
	for i := range s {
		m[s[i]-'a']++
	}

	var cnt = len(s)
	var odd = 0
	for _, v := range m {
		if v%2 == 1 {
			odd++
		}
	}

	if 0 < odd {
		odd -= 1
		cnt -= odd
	}
	return cnt
}
