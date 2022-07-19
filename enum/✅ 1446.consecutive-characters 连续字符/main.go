package main

// https://leetcode.cn/problems/consecutive-characters

func maxPower(s string) int {
	l := len(s)
	var cnt = 1
	var max int
	for i := 1; i < l; i++ {
		if s[i-1] == s[i] {
			cnt++
		} else {
			if max < cnt {
				max = cnt
			}
			cnt = 1
		}
	}
	if max < cnt {
		max = cnt
	}
	return max
}
