package main

// https://leetcode.cn/problems/letter-case-permutation

func letterCasePermutation(s string) []string {
	sL := len(s)
	var strs []string
	var buf = make([]byte, sL)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == sL {
			strs = append(strs, string(buf))
			return
		}
		if s[idx] <= '9' {
			buf[idx] = s[idx]
			dfs(idx + 1)
		} else if 'a' <= s[idx] && s[idx] <= 'z' {
			buf[idx] = s[idx]
			dfs(idx + 1)
			buf[idx] = s[idx] - 32
			dfs(idx + 1)
		} else if 'A' <= s[idx] && s[idx] <= 'Z' {
			buf[idx] = s[idx]
			dfs(idx + 1)
			buf[idx] = s[idx] + 32
			dfs(idx + 1)
		}
	}
	dfs(0)
	return strs
}
