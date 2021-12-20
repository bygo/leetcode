package main

// https://leetcode-cn.com/problems/delete-operation-for-two-strings

func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i, s1 := range word1 {
		for j, s2 := range word2 {
			if s1 == s2 {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return m + n - f[m][n]*2
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
