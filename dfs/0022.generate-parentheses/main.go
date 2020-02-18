package main

//dfs
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n < 1 {
		return res
	}
	dfs("", n, n, &res)
	return res
}

func dfs(str string, left int, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, str)
	}

	if left > right { //剪枝
		return
	}
	if left > 0 {
		dfs(str+"(", left-1, right, res) //左树遍历
	}

	if right > 0 {
		dfs(str+")", left, right-1, res) //右树遍历
	}
}

//闭合数
func generateParenthesis3(n int) []string {
	return closed(n)
}

func closed(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			leftCombination := closed(i)
			rightCombination := closed(n - 1 - i)
			for _, left := range leftCombination {
				for _, right := range rightCombination {
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
