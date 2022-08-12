package main

// https://leetcode.cn/problems/reverse-only-letters

// ❓ 仅仅反转字母

func reverseOnlyLetters(s string) string {
	var buf = []byte(s)
	var idxL, idxR = 0, len(s) - 1
	for {
		for idxL < idxR && !('a' <= s[idxL] && s[idxL] <= 'z' || 'A' <= s[idxL] && s[idxL] <= 'Z') {
			idxL++
		}

		for idxL < idxR && !('a' <= s[idxR] && s[idxR] <= 'z' || 'A' <= s[idxR] && s[idxR] <= 'Z') {
			idxR--
		}

		if idxR <= idxL {
			break
		}
		buf[idxL], buf[idxR] = buf[idxR], buf[idxL]
		idxL++
		idxR--
	}
	return string(buf)
}
