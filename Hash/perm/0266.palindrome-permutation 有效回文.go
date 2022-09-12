package main

// https://leetcode.cn/problems/palindrome-permutation

// ❓ 有效回文

func canPermutePalindrome(s string) bool {
	chMpCnt := map[byte]int{}
	for i := range s {
		chMpCnt[s[i]]++
	}

	// 奇数
	var odd = 0
	for _, cnt := range chMpCnt {
		if cnt%2 == 1 {
			odd++
		}
	}

	return odd <= 1
}
