package main

// https://leetcode-cn.com/problems/decode-ways

// 一维
// f(n) = f(n-1) + f(n-2)
func numDecodings(s string) int {
	sL := len(s)
	f := make([]int, sL+1)
	f[0] = 1
	// 0 没有路径
	if s[0] != '0' {
		f[1] = 1
	}

	// 123
	// 1的路径数 = 1
	// [1]

	// 2的路径数 = 2
	// [1] [2]
	// [12]

	// 3的路径数
	// 如果独立,即
	// [1] [2]
	// [12]

	// 如果组合,即
	// [1] [23]

	for i := 2; i <= sL; i++ {
		//
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		// 跟前置构成合法
		if s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			f[i] += f[i-2]
		}
		//if f[i] == 0 && f[i-1] == 0 {
		//	return 0
		//}
	}

	return f[sL]
}
