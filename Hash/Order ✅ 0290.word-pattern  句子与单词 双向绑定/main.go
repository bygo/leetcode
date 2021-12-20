package main

import "strings"

// https://leetcode-cn.com/problems/word-pattern

// ❓ 句子与单词 双向绑定

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	var strMpCh = map[string]byte{}
	var chMpStr = map[byte]string{}
	for i := range pattern {
		ch := strMpCh[words[i]]
		str := chMpStr[pattern[i]]
		if ch != 0 && ch != pattern[i] || str != "" && str != words[i] {
			return false
		}
		strMpCh[words[i]] = pattern[i]
		chMpStr[pattern[i]] = words[i]

	}
	return true
}
