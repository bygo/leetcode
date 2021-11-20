package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := map[byte]int{}
	var cur, max, cnt, l int
	for r := range s {
		if m[s[r]] == 0 {
			cnt++
		}
		m[s[r]]++
		for k < cnt {
			m[s[l]]--
			if m[s[l]] == 0 {
				cnt--
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
