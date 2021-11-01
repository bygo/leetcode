package main

import "strings"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters

func longestSubstring(s string, k int) int {
	var res int
	for typ := 1; typ <= 26; typ++ {
		cnt := [26]int{}
		typCnt := 0
		typValid := k
		l := 0
		for r := range s {
			ch := s[r] - 'a'
			if cnt[ch] == 0 {
				typCnt++
				typValid--
			}
			cnt[ch]++
			if cnt[ch] == k {
				typValid++
			}

			for typ < typCnt {
				ch := s[l] - 'a'
				if cnt[ch] == k {
					typValid--
				}
				cnt[ch]--
				if cnt[ch] == 0 {
					typCnt--
					typValid++
				}
				l++
			}

			if typValid == k {
				cur := r - l + 1
				if res < cur {
					res = cur
				}
			}
		}
	}
	return res
}

func longestSubstringDivide(s string, k int) int {
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
		cur := longestSubstringDivide(subStr, k)
		if res < cur {
			res = cur
		}
	}
	return res
}
