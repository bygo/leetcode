package main

// https://leetcode-cn.com/problems/shortest-palindrome

const base = 131

func shortestPalindrome(s string) string {
	n := len(s)
	var l, r, mul uint = 0, 0, 1
	index := -1
	for i := 0; i < n; i++ {
		num := uint(s[i] - '0')
		l = l*base + num
		r = r + num*mul
		if l == r {
			index = i
		}
		mul = mul * base
	}
	var res = []byte(s[index+1:])
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	res = append(res, s...)
	return string(res)
}
