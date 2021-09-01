package main

// https://leetcode-cn.com/problems/second-largest-digit-in-a-string

func secondHighest(s string) int {
	m := [10]int{}
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			m[s[i]-'0']++
		}
	}
	var cnt int
	for i := 9; 0 <= i; i-- {
		if 0 < m[i] {
			cnt++
		}
		if cnt == 2 {
			return i
		}
	}
	return -1
}
