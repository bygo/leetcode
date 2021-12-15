package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/unique-word-abbreviation

// ❓ 单词是唯一的缩写

type ValidWordAbbr struct {
	valMpComMpCnt map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	valMpComMpCnt := map[string]map[string]int{}
	for i := range dictionary {
		s := Compress(dictionary[i])
		if valMpComMpCnt[s] == nil {
			valMpComMpCnt[s] = map[string]int{}
		}
		valMpComMpCnt[s][dictionary[i]] ++
	}
	return ValidWordAbbr{
		valMpComMpCnt: valMpComMpCnt,
	}
}

func (vwa *ValidWordAbbr) IsUnique(word string) bool {
	s := Compress(word)
	return vwa.valMpComMpCnt[s] == nil || len(vwa.valMpComMpCnt[s]) == 1 && 0 < vwa.valMpComMpCnt[s][word]
}

// hello => h3o
func Compress(word string) string {
	l := len(word)
	return string(word[0]) + strconv.Itoa(l-2) + string(word[l-1])
}
