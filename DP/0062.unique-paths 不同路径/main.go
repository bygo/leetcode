package main

// https://leetcode-cn.com/problems/unique-paths

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePaths(l1, l2 int) int {
	f := make([][]int, l1)
	for i := range f {
		f[i] = make([]int, l2)
		f[i][0] = 1
	}
	for j := 0; j < l2; j++ {
		f[0][j] = 1
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[l1-1][l2-1]
}

// 压缩
func uniquePaths(l1, l2 int) int {
	f := make([]int, l2)
	f[0] = 1
	for i := 0; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[j] += f[j-1]
		}
	}
	return f[l2-1]
}
