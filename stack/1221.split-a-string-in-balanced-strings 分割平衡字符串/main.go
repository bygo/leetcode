package main

// https://leetcode.cn/problems/split-a-string-in-balanced-strings

func balancedStringSplit(s string) int {
	var cnt int
	var res int
	for i := range s {
		if s[i] == 'L' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			res++
		}
	}
	return res
}
