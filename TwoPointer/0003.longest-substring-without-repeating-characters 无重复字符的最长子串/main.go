package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

// :34
func lengthOfLongestSubstring(s string) int {
	var l, cur, res int
	for r, v := range s {
		for i := l; i < r; i++ {
			if byte(v) == s[i] {
				l = i + 1
				break
			}
		}
		cur = r - l + 1
		if res < cur {
			res = cur
		}
	}
	return res
}
