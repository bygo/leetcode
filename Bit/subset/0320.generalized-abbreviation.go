package main

import "strconv"

// https://leetcode.cn/problems/generalized-abbreviation

func generateAbbreviations(word string) []string {
	wordL := len(word)
	var subsetMax = 1 << wordL
	var strsSubset []string
	for subset := 0; subset < subsetMax; subset++ {
		var buf []byte
		var cnt int
		for idx := 0; idx < wordL; idx++ {
			if subset>>idx&1 == 1 { // or 0
				if 0 < cnt {
					buf = append(buf, strconv.Itoa(cnt)...)
				}
				cnt = 0
				buf = append(buf, word[idx])
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
