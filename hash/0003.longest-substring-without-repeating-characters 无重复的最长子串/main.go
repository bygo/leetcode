package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var res, cur, l int
	var m = map[byte]int{}
	for i := range s {
		if l < m[s[i]] {
			l = m[s[i]]
		}
		cur = i - l + 1
		if res < cur {
			res = cur
		}
		m[s[i]] = i + 1
		// 出现位置 +1
		// A.否则 [0] 重复 0 < 0 不成立
		// B.否则 cur = i - l + 1 不成立
		// C.也可 l = m[s[i]] + 1 ，但 A 不成立
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	var res, cur, l int
	var m = [127]int{}
	for i := range s {
		if l < m[s[i]] {
			l = m[s[i]]
		}

		cur = i - l + 1
		if res < cur {
			res = cur
		}
		m[s[i]] = i + 1
	}
	return res
}
