package main

// https://leetcode.cn/problems/is-unique-lcci/

func isUnique(astr string) bool {
	m := map[byte]int{}
	for i := range astr {
		m[astr[i]]++
		if 1 < m[astr[i]] {
			return false
		}
	}
	return true
}
