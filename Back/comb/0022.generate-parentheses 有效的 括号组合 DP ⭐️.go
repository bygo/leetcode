package main

// https://leetcode-cn.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	var f = [17][]string{0: {""}, 1: {"()"}}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, s1 := range f[j] {
				for _, s2 := range f[i-j-1] {
					f[i] = append(f[i], "("+s1+")"+s2)
				}
			}
		}
	}
	return f[n]
}
