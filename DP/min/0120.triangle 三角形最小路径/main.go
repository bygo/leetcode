package main

// https://leetcode.cn/problems/triangle

// 压缩 从底至顶
func minimumTotal(triangle [][]int) int {
	tL := len(triangle)
	for i := tL - 2; 0 <= i; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

// 压缩 从顶至底
func _(triangle [][]int) int {
	tL := len(triangle)
	for i := 1; i < tL; i++ {
		triangle[i][i] += triangle[i-1][i-1]
		for j := i - 1; 1 <= j; j-- {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
		triangle[i][0] += triangle[i-1][0]
	}

	for i := 1; i < tL; i++ {
		if triangle[tL-1][i] < triangle[tL-1][0] {
			triangle[tL-1][0] = triangle[tL-1][i]
		}
	}
	return triangle[tL-1][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
