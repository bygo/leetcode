package main

// https://leetcode-cn.com/problems/generate-parentheses/

func main() {
	generateParenthesis(16)
}
func generateParenthesis(n int) []string {
	var dfs func(idx int) []string
	var f = [17][]string{0: {""}}
	dfs = func(idx int) []string {
		if 0 < len(f[idx]) {
			return f[idx]
		}
		var strs []string
		for j := 0; j < idx; j++ {
			strsLeft := dfs(j)
			strsRight := dfs(idx - j - 1)
			for _, strLeft := range strsLeft {
				for _, strRight := range strsRight {
					strs = append(strs, "("+strLeft+")"+strRight)
				}
			}
			// ff = "("+strLeft+")"
			// f = strRight

			// 0 15
			// ""
			// 1 = ff(0) * f(0)
			// ()
			// 2 = ff(1) * f(0) + ff(0) * f(1)
			// ()()  (())
			// 3 = ff(0) * f(2) + ff(1) * f(1) + ff(2) * f(0)
			// ((())) (()()) (())() ()(()) ()()()
			// 4 = 3*0 2*1 1*2 0*3
			// 5 = 4*0 3*1 2*2 1*3 0*4
			// 6 = 5*0 4*1 3*2 2*3 1*4 0*5
		}
		f[idx] = strs
		return f[idx]
	}
	return dfs(n)
}
