package main

// https://leetcode-cn.com/problems/palindrome-permutation-lcci/

func canPermutePalindrome(s string) bool {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}
	var cnt int
	for b := range m {
		if m[b]%2 == 1 {
			cnt++
		}
	}
	return cnt <= 1
}
