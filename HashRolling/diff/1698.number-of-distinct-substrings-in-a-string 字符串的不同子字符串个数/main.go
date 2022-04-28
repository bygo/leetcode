package main

// https://leetcode-cn.com/problems/number-of-distinct-substrings-in-a-string

// ❓ 字符串的不同子字符串个数

const base = 131

func countDistinct(s string) int {
	sL := len(s)

	hashMp := map[uint]struct{}{}
	var hash uint
	for idx := 0; idx < sL; idx++ {
		maxL := sL - idx
		hash = 0
		for curL := 0; curL < maxL; curL++ {
			hash = hash*base + uint(s[idx+curL])
			hashMp[hash] = struct{}{}
		}
	}
	return len(hashMp)
}
