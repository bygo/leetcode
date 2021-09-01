package main

import "strings"

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")
	m := map[byte]bool{}
	for i := range brokenLetters {
		m[brokenLetters[i]] = true
	}

	var res int
	for i := range words {
		j := 0
		for j < len(words[i]) {
			if m[words[i][j]] {
				break
			}
			j++
		}
		if j == len(words[i]) {
			res++
		}
	}
	return res
}
