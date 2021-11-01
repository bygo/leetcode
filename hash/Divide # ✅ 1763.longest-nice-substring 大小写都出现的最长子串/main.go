package main

// https://leetcode-cn.com/problems/longest-nice-substring

func longestNiceSubstring(s string) string {
	var dfs func(sub string) string
	dfs = func(sub string) string {
		var m = map[byte]int{}
		for i := range sub {
			if sub[i] <= 'Z' && (m[sub[i]+32] == 0 || m[sub[i]+32] == 1) {
				m[sub[i]+32] += 2
			} else if 'a' <= sub[i] && (m[sub[i]] == 0 || m[sub[i]] == 2) {
				m[sub[i]] += 1
			}
		}

		var left int
		var res string
		for j := range sub {
			ch := sub[j]
			if ch <= 'Z' {
				ch += 32
			}
			if m[ch] != 3 {
				cur := dfs(sub[left:j])
				if len(res) < len(cur) {
					res = cur
				}
				left = j + 1
			}
		}
		if 0 < left {
			cur := dfs(sub[left:])
			if len(res) < len(cur) {
				res = cur
			}
			return res
		}
		return sub
	}
	return dfs(s)
}
