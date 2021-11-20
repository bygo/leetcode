package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := [128]int{}
	var n = len(s)
	var max, cnt, l, r int
	var f = func() {
		cur := r - l
		if max < cur {
			max = cur
		}
	}
	for r < n {
		if m[s[r]] == 0 {
			f()
			cnt++
			for 2 < cnt {
				m[s[l]]--
				if m[s[l]] == 0 {
					cnt--
				}
				l++
			}
		}
		m[s[r]]++
		r++
	}
	f()

	return max
}
