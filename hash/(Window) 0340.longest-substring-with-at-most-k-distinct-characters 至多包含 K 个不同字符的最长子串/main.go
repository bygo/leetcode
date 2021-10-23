package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := map[byte]int{}
	var cur, max, l int
	for r := range s {
		m[s[r]]++
		for k < len(m) {
			m[s[l]]--
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}
		cur = r - l + 1
		if max < cur {
			max = cur
		}
	}
	return max
}
