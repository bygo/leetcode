package main

// https://leetcode-cn.com/problems/longest-palindromic-substring

// 一维
func longestPalindrome(s string) string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n) // f[left][right] left至right 是否回文
	}

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || f[l+1][r-1]) { // 0:相等  1:相邻  2:隔一个字符
				f[l][r] = true
				if right-left < r-l {
					right = r
					left = l
				}
			}
		}
	}
	return s[left : right+1]
}

// 压缩外循环
func longestPalindrome(s string) string {
	n := len(s)
	f := make([]bool, n)

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			f[l] = s[l] == s[r] && (r-l <= 2 || f[l+1])
			if f[l] && right-left < r-l {
				right = r
				left = l
			}
		}
	}
	return s[left : right+1]
}
