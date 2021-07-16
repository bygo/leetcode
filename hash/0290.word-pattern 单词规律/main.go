package main

import "strings"

// https://leetcode-cn.com/problems/word-pattern

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	var patternSet = map[string]byte{}
	var wordsSet = map[byte]string{}
	for i := range pattern {
		p := patternSet[words[i]]
		w := wordsSet[pattern[i]]
		if 0 != p && p != pattern[i] || 0 < len(w) && w != words[i] {
			return false
		}
		patternSet[words[i]] = pattern[i]
		wordsSet[pattern[i]] = words[i]
	}
	return true
}
