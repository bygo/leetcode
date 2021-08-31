package main

// https://leetcode-cn.com/problems/palindrome-permutation

func canPermutePalindrome(s string) bool {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	var c = 0
	for _, v := range m {
		if v%2 == 1 {
			c++
		}
	}

	return c <= 1
}
