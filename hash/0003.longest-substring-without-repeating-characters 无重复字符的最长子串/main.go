package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var l, cur, res int
	var m = [127]int{}
	for i := range s {
		if l < m[s[i]] {
			l = m[s[i]]
		}

		m[s[i]] = i + 1
		cur = i - l + 1
		if res < cur {
			res = cur
		}
	}
	return res
}
