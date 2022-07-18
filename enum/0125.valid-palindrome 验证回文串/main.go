package main

// https://leetcode.cn/problems/valid-palindrome

func isPalindrome(s string) bool {
	l1 := len(s)
	b := make([]byte, 0, l1)
	for i := 0; i < l1; i++ {
		ch := s[i]
		if 'a' <= ch && ch <= 'z' || '0' <= ch && ch <= '9' {
			b = append(b, ch)
		} else if 'A' <= ch && ch <= 'Z' {
			b = append(b, ch+32)
		}
	}

	l2 := len(b)
	for i := 0; i < l2/2; i++ {
		if b[i] != b[l2-i-1] {
			return false
		}
	}
	return true
}
