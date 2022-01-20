package main

// https://leetcode-cn.com/problems/palindrome-permutation

// ❓ 有效回文

func canPermutePalindrome(s string) bool {
	var bit int
	for i := range s {
		bit ^= 1 << (s[i] - 'a')
	}
	return bit&(bit-1) == 0
}
