package main

// https://leetcode.cn/problems/longest-repeating-character-replacement

func characterReplacement(s string, k int) int {
	n := len(s)
	l := 0
	m := map[byte]int{}
	var max int
	for r := range s {
		m[s[r]]++
		cur := m[s[r]]
		if max < cur {
			max = cur
		}
		if k < r-l+1-max {
			m[s[l]]--
			l++
		}
	}
	return n - l
}
