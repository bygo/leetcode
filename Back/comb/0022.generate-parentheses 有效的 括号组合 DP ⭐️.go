package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var f = [9][]string{{""}, {"()"}}
	for i := 2; i <= n; i++ { // 括号长度
		for j := 0; j < i; j++ { // 总组合数
			for _, strLeft := range f[j] { // 左边
				for _, strRight := range f[i-j-1] { // 右边
					f[i] = append(f[i], "("+strLeft+")"+strRight)
				}
			}
		}
	}
	return f[n]
}
