package main

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
 */

//深度搜索
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

//广度搜索
func generateParenthesis2(n int) []string {
	return nil
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
			for _, left := range closed(i) {
				for _, right := range closed(n - 1 - i) {
					res = append(res, "("+left+")"+right)
				}
			}
		}
	}
	return res
}
