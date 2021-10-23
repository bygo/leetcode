package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := map[byte]int{}
	var left int
	var max int
	for i := range s {
		m[s[i]] = i
		if len(m) < 3 {
			cur := i - left + 1
			if max < cur {
				max = cur
			}
		} else {
			var min = 1<<63 - 1
			for j := range m {
				if m[j] < min {
					min = m[j]
				}
			}
			k := s[min]
			left = m[k] + 1
			delete(m, k)
		}
	}

	cur := len(s) - left + 1
	if max < cur {
		max = cur
	}

	return max
}
