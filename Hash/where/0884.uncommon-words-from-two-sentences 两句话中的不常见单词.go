package main

import "strings"

// https://leetcode.cn/problems/uncommon-words-from-two-sentences

// ❓ 在两个数组只出现一次的单词

func uncommonFromSentences(s1 string, s2 string) []string {
	// 计数
	strings1 := strings.Split(s1, " ")
	strings2 := strings.Split(s2, " ")

	strMpCnt := map[string]int{}

	for _, str := range strings1 {
		strMpCnt[str]++
	}

	for _, str := range strings2 {
		strMpCnt[str]++
	}

	// 是否一次
	var strsUnique []string
	for str := range strMpCnt {
		if strMpCnt[str] == 1 {
			strsUnique = append(strsUnique, str)
		}
	}
	return strsUnique
}
