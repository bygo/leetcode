package main

// https://leetcode-cn.com/problems/distinct-subsequences

// 选择或不选择

func numDistinct(s string, t string) int {
	sL, tL := len(s), len(t)
	f := make([]int, tL+1)
	f[0] = 1

	for i := 1; i <= sL; i++ {
		angle := f[0]
		for j := 1; j <= tL; j++ {
			// 不选择当前字符 的 路径数
			pre := f[j]
			// 选择当前字符 的 路径数 ，必须相等
			if s[i-1] == t[j-1] {
				// 相当于消元
				f[j] += angle
			}

			angle = pre
		}
	}

	return f[tL]
}
