package main

func generateParenthesis(n int) []string {
	var res []string
	dfs(nil, n, n, &res)
	return res
}

func dfs(cur []byte, l int, r int, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, string(cur))
	}

	if l <= r {
		if 0 < l {
			dfs(append(cur, '('), l-1, r, res)
		}

		if 0 < r {
			dfs(append(cur, ')'), l, r-1, res)
		}
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
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
