package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var buf = make([]byte, n*2)
	var strs []string
	var n8 = int8(n)
	var cntL, cntR int8
	var dfs func()
	dfs = func() {
		// 左边 <= 右边 永远合法
		if cntL < cntR {
			return
		}
		if cntR == n8 && cntL == n8 {
			strs = append(strs, string(buf))
			return
		}
		if cntL < n8 {
			// 添加一个左括号
			buf[cntL+cntR] = '('
			cntL++
			dfs()
			cntL--
		}
		// 添加一个右括号
		buf[cntL+cntR] = ')'
		cntR++
		dfs()
		cntR--
	}
	dfs()
	return strs
}
