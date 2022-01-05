package main

// https://leetcode-cn.com/problems/longest-happy-prefix

// ❓ 最长快乐前缀

const base uint = 131

func longestPrefix(s string) string {
	// 自身不参与计算
	sTop := len(s) - 1
	var hashLeft, hashRight, mul uint = 0, 0, 1
	var idx = -1
	for i := 0; i < sTop; i++ {
		bitLeftLo := uint(s[i])
		bitRightHi := uint(s[sTop-i])
		hashLeft = hashLeft*base + bitLeftLo
		hashRight = bitRightHi*mul + hashRight
		if hashLeft == hashRight {
			idx = i
		}
		mul *= base
	}
	return s[:idx+1]
}
