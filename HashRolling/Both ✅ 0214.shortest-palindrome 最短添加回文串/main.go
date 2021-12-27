package main

// https://leetcode-cn.com/problems/shortest-palindrome

// ❓ 最短回文串

const base = 131

func shortestPalindrome(s string) string {
	sL := len(s)
	var hashLeft, hashRight, mul uint = 0, 0, 1
	var idx = -1
	for i := 0; i < sL; i++ {
		bitLow := uint(s[i])
		hashLeft = hashLeft*base + bitLow
		hashRight = hashRight + bitLow*mul
		if hashLeft == hashRight {
			idx = i
		}
		mul *= base
	}

	strNeed := []byte(s[idx+1:])
	left, right := 0, len(strNeed)-1
	for left < right {
		strNeed[left], strNeed[right] = strNeed[right], strNeed[left]
		left++
		right--
	}

	bufPalindrome := append(strNeed, s...)
	return string(bufPalindrome)
}
