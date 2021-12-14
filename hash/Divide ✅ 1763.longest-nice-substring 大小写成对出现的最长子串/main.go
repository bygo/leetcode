package main

// https://leetcode-cn.com/problems/longest-nice-substring

// ❓大小写成对出现的最长子串

func longestNiceSubstring(s string) string {
	byteMp := map[byte]struct{}{}
	for i := range s {
		byteMp[s[i]] = struct{}{}
	}

	for i := range s {
		// 找对
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else {
			ch -= 32
		}
		_, ok := byteMp[ch]

		if !ok {
			// 分区
			left := longestNiceSubstring(s[:i])
			right := longestNiceSubstring(s[i+1:])
			if len(left) < len(right) {
				return right
			}
			return left
		}
	}
	return s
}
