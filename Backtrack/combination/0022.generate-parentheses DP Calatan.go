package main

// https://leetcode.cn/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var f = [9][]string{{""}, {"()"}, {"()()", "(())"}}
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, strLeft := range f[j] {
				for _, strRight := range f[i-j-1] {
					f[i] = append(f[i], "("+strLeft+")"+strRight)
				}
			}
		}
	}
	return f[n]
}
