package main

// Link: https://leetcode-cn.com/problems/interleaving-string

// 二维
// f(i)(j) = f(i)(j-1) && s2[j-1] == p || f(i-1)(j) && s1[i-1] == p
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	f := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]bool, l2+1)
	}
	f[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			f[i][j] = i == 0 && j == 0 || 0 < i && f[i-1][j] && s1[i-1] == s3[i+j-1] ||
				0 < j && f[i][j-1] && s2[j-1] == s3[i+j-1]
		}
	}
	return f[l1][l2]
}

// 压缩
// dp[j] 表示 当前i下
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	f := make([]bool, l2+1)
	f[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if 0 < i {
				f[j] = f[j] && s1[i-1] == s3[i+j-1] // i增加后 f[j] 是否依然有效
			}

			if 0 < j { // 如果dp[j] 有效 或 可以通过增加j 使之有效
				f[j] = f[j] || f[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return f[l2]
}
