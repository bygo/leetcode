package main

// https://leetcode-cn.com/problems/generate-parentheses/

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
