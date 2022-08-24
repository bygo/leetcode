package main

// https://leetcode.cn/problems/palindrome-permutation

func canPermutePalindrome(s string) bool {
	var bit int
	for i := range s {
		bit ^= 1 << (s[i] - 'a') // TODO
	}
	return bit&(bit-1) == 0
}
