package main

// https://leetcode.cn/problems/decode-ways

// f(n) = f(n-1) + f(n-2)

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

		if s[i-1] != '0' {
			num := (s[i-1]-'0')*10 + s[i] - '0'
			if num <= 26 {
				z += x
			}
		}

		x, y = y, z
	}
	return z
}
