package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var buf = make([]byte, n*2)
	var strs []string
	var nInt8 = int8(n)
	var cntL, cntR int8
	var dfs func()
	dfs = func() {
		if cntL < cntR {
			return
		}
		if cntR == nInt8 && cntL == nInt8 {
			strs = append(strs, string(buf))
			return
		}

		if cntL < nInt8 {
			buf[cntL+cntR] = '('
			cntL++
			dfs()
			cntL--
		}
		buf[cntL+cntR] = ')'
		cntR++
		dfs()
		cntR--
	}
	dfs()
	return strs
}
