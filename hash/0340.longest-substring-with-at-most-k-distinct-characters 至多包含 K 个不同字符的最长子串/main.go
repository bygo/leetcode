package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := map[byte]int{}
	var cur, max, left int
	for right := range s {
		m[s[right]]++
		for k < len(m) {
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
			left++
		}
		cur = right - left + 1
		if max < cur {
			max = cur
		}
	}
	return max
}
