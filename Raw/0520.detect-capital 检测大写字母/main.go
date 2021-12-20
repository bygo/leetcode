package main

// https://leetcode-cn.com/problems/detect-capital

func detectCapitalUse(word string) bool {
	l := len(word)
	if l == 1 {
		return true
	}
	var left, right byte = 'a', 'z'
	if 'A' <= word[0] && word[0] <= 'Z' && 'A' <= word[1] && word[1] <= 'Z' {
		left = 'A'
		right = 'Z'
	}

	var i = 1
	for i < l && left <= word[i] && word[i] <= right {
		i++
	}
	return i == l
}
