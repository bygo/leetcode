package main

// https://leetcode-cn.com/problems/longest-happy-prefix

func longestPrefix(s string) string {
	l := len(s)
	var p = -1
	for i := 1; i < l; i++ {
		if s[:i] == s[l-i:] {
			p = i
		}
	}
	if -1 < p {
		return s[:p]
	}
	return ""
}

const base uint = 131

func longestPrefix(s string) string {
	n := len(s) - 1
	var l, r uint
	var mul uint = 1
	var p = -1
	for i := 0; i < n; i++ {
		l = l*base + uint(s[i])
		r = uint(s[n-i])*mul + r
		if l == r {
			p = i
		}
		mul *= base
	}
	return s[:p]
}
