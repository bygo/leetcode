package main

// https://leetcode-cn.com/problems/shortest-palindrome

// ❓ 最短回文串

const base = 131

func shortestPalindrome(s string) string {
	sL := len(s)
	var hashLeft, hashRight, mul uint = 0, 0, 1
	idx := -1
	for i := 0; i < sL; i++ {
		bit := uint(s[i])
		hashLeft = hashLeft*base + bit
		hashRight = hashRight + bit*mul
		// 寻找最长回文
		if hashLeft == hashRight {
			idx = i
		}
		// 随着长度增加变大
		mul = mul * base
	}

	// 需要添加的字符
	var strNeed = []byte(s[idx+1:])

	// 反转
	i, j := 0, len(strNeed)-1
	for i < j {
		strNeed[i], strNeed[j] = strNeed[j], strNeed[i]
		i++
		j--
	}
	strPalindrome := append(strNeed, s...)
	return string(strPalindrome)
}

func shortestPalindrome(s string) string {
	sL := len(s)

}
