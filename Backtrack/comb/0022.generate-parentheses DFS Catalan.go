package main

// https://leetcode.cn/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var dfs func(idx int) []string
	var cache = [9][]string{0: {""}}
	dfs = func(idx int) []string {
		if 0 < len(cache[idx]) {
			return cache[idx]
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
		}
		cache[idx] = strs
		return cache[idx]
	}
	return dfs(n)
}

// ff = "("+strLeft+")"
// f = strRight

// 0 15
// ""
// f(1) = ff(0) * f(0)
// ()
// f(2) = ff(1) * f(0) + ff(0) * f(1)
// ()()  (())
// f(3) = ff(0) * f(2) + ff(1) * f(1) + ff(2) * f(0)
// ((())) (()()) (())() ()(()) ()()()

// f(4) = 3*0 2*1 1*2 0*3
// f(5) = 4*0 3*1 2*2 1*3 0*4
// f(6) = 5*0 4*1 3*2 2*3 1*4 0*5
