package main

import "strings"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters

func longestSubstring(s string, k int) int {
	var res int
	if s == "" {
		return res
	}

	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		cur := longestSubstring(subStr, k)
		if res < cur {
			res = cur
		}
	}
	return res
}
