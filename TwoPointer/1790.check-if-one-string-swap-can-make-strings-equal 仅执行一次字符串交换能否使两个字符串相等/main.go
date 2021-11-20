package main

// https://leetcode-cn.com/problems/check-if-one-string-swap-can-make-strings-equal

func areAlmostEqual(s1 string, s2 string) bool {
	var first, second, flag byte
	for k := range s1 {
		if s2[k] != s1[k] {
			flag++
			if first == 0 {
				first, second = s1[k], s2[k]
			} else if first != s2[k] || second != s1[k] || 2 < flag {
				return false
			}
		}
	}
	return flag != 1
}
