package main

import (
	"strconv"
)

// https://leetcode.cn/problems/unique-word-abbreviation

// ❓ 单词是唯一的缩写

type ValidWordAbbr struct {
	comMpStrMpCnt map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	comMpStrMpCnt := map[string]map[string]int{}
	for _, str := range dictionary {
		com := Compress(str)
		if comMpStrMpCnt[com] == nil {
			comMpStrMpCnt[com] = map[string]int{}
		}
		comMpStrMpCnt[com][str]++
	}
	return ValidWordAbbr{
		comMpStrMpCnt: comMpStrMpCnt,
	}
}

func (vwa *ValidWordAbbr) IsUnique(str string) bool {
	com := Compress(str)
	return vwa.comMpStrMpCnt[com] == nil || len(vwa.comMpStrMpCnt[com]) == 1 && 0 < vwa.comMpStrMpCnt[com][str]
}

// hello => h3o

func Compress(str string) string {
	wordL := len(str)
	return string(str[0]) + strconv.Itoa(wordL-2) + string(str[wordL-1])
}
