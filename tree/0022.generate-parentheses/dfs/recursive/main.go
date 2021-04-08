package main

func generateParenthesis(n int) []string {
	var res []string
	dfs("", n, n, &res)
	return res
}

func dfs(cur string, left int, right int, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}

	if left <= right {
		if 0 < left {
			dfs(cur+"(", left-1, right, res)
		}

		if 0 < right {
			dfs(cur+")", left, right-1, res)
		}
	}
}
