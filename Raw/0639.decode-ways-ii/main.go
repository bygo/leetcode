package main

// https://leetcode-cn.com/problems/decode-ways-ii

func check1digit(ch byte) int {
	if ch == '*' {
		return 9
	}
	if ch == '0' {
		return 0
	}
	return 1
}

func check2digits(c0, c1 byte) int {
	if c0 == '*' && c1 == '*' {
		return 15
	}
	if c0 == '*' {
		if c1 <= '6' {
			return 2
		}
		return 1
	}
	if c1 == '*' {
		if c0 == '1' {
			return 9
		}
		if c0 == '2' {
			return 6
		}
		return 0
	}
	if c0 != '0' && (c0-'0')*10+(c1-'0') <= 26 {
		return 1
	}
	return 0
}

func numDecodings(s string) int {
	const mod int = 1e9 + 7
	a, b, c := 0, 1, 0
	for i := range s {
		c = b * check1digit(s[i]) % mod
		if 0 < i {
			c = (c + a*check2digits(s[i-1], s[i])) % mod
		}
		a, b = b, c
	}
	return c
}
