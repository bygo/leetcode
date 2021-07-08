package main

// Link: https://leetcode-cn.com/problems/triangle

// 压缩 从顶至底
func minimumTotal(triangle [][]int) int {
	l1 := len(triangle)
	l2 := len(triangle[l1-1])
	var f = make([]int, l2)
	f[0] = triangle[0][0]

	for i := 1; i < l1; i++ {
		f[i] = f[i-1] + triangle[i][i] // 最后一个
		for j := i - 1; 0 < j; j-- {   // 中间
			f[j] = min(f[j], f[j-1]) + triangle[i][j]
		}
		f[0] = f[0] + triangle[i][0] // 第一个
	}

	for i := 1; i < l2; i++ {
		if f[i] < f[0] {
			f[0] = f[i]
		}
	}

	return f[0]
}

// 压缩 从底至顶
func minimumTotal(triangle [][]int) int {
	l1 := len(triangle)
	l2 := len(triangle[l1-1])
	var f = make([]int, l2+1)
	for i := l2 - 1; 0 <= i; i-- {
		for j := 0; j <= i; j++ {
			f[j] = min(f[j], f[j+1]) + triangle[i][j]
		}
	}

	return f[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
