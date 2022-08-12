package main

// https://leetcode.cn/problems/longest-palindrome

// ❓ 能排列为最长的回文串

func longestPalindrome(s string) int {
	chMpCnt := map[byte]int{}
	for i := range s {
		ch := s[i] - 'a'
		chMpCnt[ch]++
	}

	var cntLongest = len(s)
	var odd = 0
	for _, cnt := range chMpCnt {
		if cnt%2 == 1 {
			odd++
		}
	}

	if 0 < odd {
		return cntLongest - odd + 1
	}
	return cntLongest
}
