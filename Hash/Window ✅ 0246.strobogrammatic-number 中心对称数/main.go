package main

// https://leetcode-cn.com/problems/strobogrammatic-number

func isStrobogrammatic(num string) bool {
	m := map[byte]byte{'6': '9', '9': '6', '8': '8', '0': '0', '1': '1'}

	l, r := 0, len(num)-1
	for l <= r {
		if m[num[l]] != num[r] || m[num[l]] == 0 {
			return false
		}
		l++
		r--
	}
	return true
}
