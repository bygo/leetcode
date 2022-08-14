package main

// https://leetcode.cn/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var cache = [9][]string{0: {""}}
	var dfs func(idx int) []string
	dfs = func(idx int) []string {
		if len(cache[idx]) == 0 {
			for i := 0; i < idx; i++ {
				strsLeft := dfs(i)
				strsRight := dfs(idx - i - 1)
				for _, strLeft := range strsLeft {
					for _, strRight := range strsRight {
						cache[idx] = append(cache[idx], "("+strLeft+")"+strRight)
					}
				}
			}
		}
		return cache[idx]
	}
	return dfs(n)
}

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
