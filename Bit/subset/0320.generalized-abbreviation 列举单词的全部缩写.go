package main

import "strconv"

// https://leetcode-cn.com/problems/generalized-abbreviation

// ❓ 列举单词的全部缩写

func generateAbbreviations(word string) []string {
	wordL := len(word)
	var numMax = 1<<wordL - 1
	var strsSubset []string
	for subset := 0; subset < numMax; subset++ {
		var buf []byte
		var cnt int
		for pos := 0; pos < wordL; pos++ {
			if subset>>pos&1 == 1 {
				if 0 < cnt {
					buf = append(buf, strconv.Itoa(cnt)...)
				}
				cnt = 0
				buf = append(buf, word[pos])
			} else {
				cnt++
			}
		}
		if 0 < cnt {
			buf = append(buf, strconv.Itoa(cnt)...)
		}
		strsSubset = append(strsSubset, string(buf))
	}
	return strsSubset
}
