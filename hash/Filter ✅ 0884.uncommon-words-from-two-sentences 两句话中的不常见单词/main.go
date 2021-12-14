package main

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences

// ❓在两个数组只出现一次的单词

func uncommonFromSentences(s1 string, s2 string) []string {
	// 统计
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")

	wordMpCnt := map[string]int{}

	for i := range a1 {
		wordMpCnt[a1[i]]++
	}

	for i := range a2 {
		wordMpCnt[a2[i]]++
	}

	// 是否只一次
	var res []string
	for word := range wordMpCnt {
		if wordMpCnt[word] == 1 {
			res = append(res, word)
		}
	}
	return res
}
