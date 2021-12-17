package main

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences

// ❓ 在两个数组只出现一次的单词

func uncommonFromSentences(s1 string, s2 string) []string {
	// 计数
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")

	strMpCnt := map[string]int{}

	for _, str := range a1 {
		strMpCnt[str]++
	}

	for _, str := range a2 {
		strMpCnt[str]++
	}

	// 是否一次
	var res []string
	for str := range strMpCnt {
		if strMpCnt[str] == 1 {
			res = append(res, str)
		}
	}
	return res
}
