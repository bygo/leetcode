package main

// https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition

func longestDecomposition(text string) int {
	l := len(text)
	if l <= 1 {
		return l
	}
	for i := 1; i < l; i++ {
		if text[:i] == text[l-i:] {
			return 2 + longestDecomposition(text[i:l-i])
		}
	}
	return 1
}
