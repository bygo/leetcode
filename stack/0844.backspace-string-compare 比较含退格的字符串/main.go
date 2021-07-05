package main

// Link: https://leetcode-cn.com/problems/backspace-string-compare

func backspaceCompare(s string, t string) bool {
	var ss, tt []byte
	for i := range s {
		if s[i] == '#' {
			if 0 < len(ss) { // 为了写清楚逻辑
				ss = ss[:len(ss)-1]
			}
		} else {
			ss = append(ss, s[i])
		}
	}

	for i := range t {
		if t[i] != '#' {
			tt = append(tt, t[i])
		} else if 0 < len(tt) {
			tt = tt[:len(tt)-1]
		}
	}

	return string(ss) == string(tt)
}
