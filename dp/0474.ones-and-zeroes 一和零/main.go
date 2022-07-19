package main

// https://leetcode.cn/problems/ones-and-zeroes

func findMaxForm(strs []string, m, n int) int {
	sL := len(strs)
	f := make([][][]int, sL+1)
	for i := range f {
		f[i] = make([][]int, m+1)
		for j := range f[i] {
			f[i][j] = make([]int, n+1)
		}
	}

	for i, s := range strs {
		zeros, ones := get(s)
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				f[i+1][j][k] = f[i][j][k]
				if zeros <= j && ones <= k {
					// 第i个
					f[i+1][j][k] = max(f[i+1][j][k], f[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return f[sL][m][n]
}

// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
func findMaxForm(strs []string, m, n int) int {
	f := make([][]int, m+1)

	for i := range f {
		f[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := get(s)

		for i := m; zeros <= i; i-- {
			for j := n; ones <= j; j-- {
				f[i][j] = max(f[i][j], f[i-zeros][j-ones]+1)
			}
		}
	}
	return f[m][n]
}

func get(s string) (zeros, ones int) {
	for _, v := range s {
		if v == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
