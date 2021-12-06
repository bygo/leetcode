package main

// https://leetcode-cn.com/problems/longest-nice-substring

func longestNiceSubstring(s string) string {
	m := map[byte]struct{}{}
	for i := range s {
		m[s[i]] = struct{}{}
	}

	for i := range s {
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else {
			ch -= 32
		}
		_, ok := m[ch]
		if !ok {
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
