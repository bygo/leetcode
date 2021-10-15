package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var res []string
	var dfs func(cur string, l, r int)
	dfs = func(cur string, l, r int) {
		if l == 0 && r == 0 {
			res = append(res, cur)
		}

		if l <= r {
			if 0 < l {
				dfs(cur+"(", l-1, r)
			}

			if 0 < r {
				dfs(cur+")", l, r-1)
			}
		}
	}
	dfs("", n, n)
	return res
}

func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range generateParenthesis(i) {
				for _, right := range generateParenthesis(n - 1 - i) {
					res = append(res, "("+left+")"+right)
				}
			}
			// （）
			//  0 = ""
			//  1 = "" ()
			//  2 = "" ()()  (())
			//  3 = "" (()())  ()(())  (())() ((()))
		}
	}
	return res
}
