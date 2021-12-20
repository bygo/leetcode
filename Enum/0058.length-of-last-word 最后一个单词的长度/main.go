package main

// https://leetcode-cn.com/problems/length-of-last-word

func lengthOfLastWord(s string) {
	var res int
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	for 0 <= i && s[i] != ' ' {
		res++
		i--
	}
	return
}
