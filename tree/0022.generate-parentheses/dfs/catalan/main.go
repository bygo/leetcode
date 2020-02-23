package main

func generateParenthesis(n int) []string {
	return dfs(n)
}

func dfs(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			for _, left := range  dfs(i) {
				for _, right := range dfs(n - 1 - i) {
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
