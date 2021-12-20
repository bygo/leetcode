package main

// https://leetcode-cn.com/problems/decode-ways

// 一维
// f(n) = f(n-1) + f(n-2)
func numDecodings(s string) int {
	l1 := len(s)
	f := make([]int, l1+1)
	f[0] = 1
	if s[0] != '0' {
		f[1] = 1
	}

	for i := 2; i <= l1; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}

		if s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			f[i] += f[i-2]
		}
	}

	return f[l1]
}

// 压缩
func numDecodings(s string) int {
	l1 := len(s)
	x, y, z := 1, 0, 0
	if s[0] != '0' {
		y = 1
		z = 1
	}

	for i := 1; i < l1; i++ {
		if s[i] != '0' {
			z = y
		} else {
			z = 0
		}

		if s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0') <= 26) {
			z += x
		}

		x, y = y, z
	}
	return z
}
