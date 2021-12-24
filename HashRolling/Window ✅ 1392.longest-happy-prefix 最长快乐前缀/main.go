package main

// https://leetcode-cn.com/problems/longest-happy-prefix

// ❓ 最长快乐前缀

const base uint = 131

func longestPrefix(s string) string {
	sTop := len(s) - 1
	var hashLeft, hashRight uint
	var mul uint = 1
	var idx = -1
	for i := 0; i < sTop; i++ {
		numLeft := uint(s[i])
		numRight := uint(s[sTop-i])

		hashLeft = hashLeft*base + numLeft
		hashRight = numRight*mul + hashRight

		// 合法快乐前缀
		if hashLeft == hashRight {
			idx = i
		}
		// 翻倍
		mul *= base
	}
	return s[:idx]
}
