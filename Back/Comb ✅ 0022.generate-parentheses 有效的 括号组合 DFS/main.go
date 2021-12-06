package main

// https://leetcode-cn.com/problems/generate-parentheses/

func main() {
	generateParenthesis(16)
}
func generateParenthesis(n int) []string {
	n2 := n * 2
	var cur = make([]byte, n2)
	var res []string
	var cnt int
	var dfs func(l, r int)
	dfs = func(l, r int) {
		cnt++

		if r < l {
			return
		}
		if l == 0 && r == 0 {
			res = append(res, string(cur))
			return
		}

		if 0 < l {
			cur[n2-l-r] = '('
			dfs(l-1, r)
		}

		if 0 < r {
			cur[n2-l-r] = ')'
			dfs(l, r-1)
		}
	}
	dfs(n, n)
	println(cnt)
	return res
}
