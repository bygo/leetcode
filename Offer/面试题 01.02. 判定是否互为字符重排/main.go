package main

// https://leetcode.cn/problems/check-permutation-lcci/

func CheckPermutation(s1 string, s2 string) bool {
	m := map[byte]int{}
	for i := range s1 {
		m[s1[i]]++
	}

	for i := range s2 {
		m[s2[i]]--
	}

	for b := range m {
		if m[b] != 0 {
			return false
		}
	}
	return true
}
