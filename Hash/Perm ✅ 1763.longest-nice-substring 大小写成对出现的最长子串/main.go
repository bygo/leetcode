package main

// https://leetcode-cn.com/problems/longest-nice-substring

// ❓ 大小写成对出现的最长子串
// ⚠️ Aa Bb... 成对

func longestNiceSubstring(s string) string {
	chMp := map[byte]struct{}{}
	for i := range s {
		chMp[s[i]] = struct{}{}
	}

	for i := range s {
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else if 'a' <= ch {
			ch -= 32
		}
		_, ok := chMp[ch]
		if !ok {
			left := longestNiceSubstring(s[:i])
			if len(s[i+1:]) <= len(left) {
				return left
			}
			right := longestNiceSubstring(s[i+1:])
			if len(left) < len(right) {
				return right
			}
			return left
		}
	}
	return s
}
