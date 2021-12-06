package main

// https://leetcode-cn.com/problems/generate-parentheses/

func main() {
	generateParenthesis(16)
}
func generateParenthesis(n int) []string {
	var dfs func(i int) []string
	var f = [17][]string{0: {""}}
	var cnt int
	dfs = func(i int) []string {
		if 0 == len(f[i]) {
			var res []string
			for j := 0; j < i; j++ {
				lefts := dfs(j)
				rights := dfs(i - j - 1)
				for _, left := range lefts {
					for _, right := range rights {
						cnt++
						res = append(res, "("+left+")"+right)
					}
				}
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
			f[i] = res
		}

		return f[i]
	}
	res := dfs(n)
	println(cnt)
	return res
}
