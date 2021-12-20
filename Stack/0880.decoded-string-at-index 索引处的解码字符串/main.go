package main

// https://leetcode-cn.com/problems/decoded-string-at-index

func decodeAtIndex(s string, k int) string {
	var size = 0
	n := len(s)

	// 递增
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			size *= int(s[i] - '0')
		} else {
			size += 1
		}
	}

	// 递减
	for i := n - 1; 0 <= i; i-- {
		k = k % size
		// 回退到指定位置时 就可以开始判断
		// 为什么不用k=size ? 因为回退时可能回退到小于size，
		if k == 0 && 'a' <= s[i] && s[i] <= 'z' {
			return string(s[i])
		}

		if '0' <= s[i] && s[i] <= '9' {
			size = size / int(s[i]-'0')
		} else {
			size -= 1
		}
	}
	return ""
}
