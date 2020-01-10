package main

import (
	"leetcode"
)

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
 */

func longestPalindrome(s string) string {
	m := "^#"
	for _, v := range s {
		m += string(v) + "#"
	}
	m += "$"
	l := len(m)
	p := make([]int, l)
	var maxCenter int
	var maxRight int
	var center int
	var right int
	for k, _ := range m {
		if k == 0 || k == l-1 {
			continue
		}
		if maxRight > l-k { //不够长，直接break
			break
		}
		p[k] = 1
		if k < right { //镜像
			mirror := p[center*2-k]
			if mirror+k > right {
				p[k] = right - k
			} else {
				p[k] = p[center*2-k]
			}
		}

		for m[k-p[k]] == m[k+p[k]] { //扩散
			p[k]++
			center = k
			right = k + p[k]
		}

		if p[k] > maxRight {
			maxCenter = k
			maxRight = p[k]
		}
	}

	var result string
	for _, v := range m[maxCenter-maxRight+1 : maxCenter+maxRight] {
		if v != '#' {
			result += string(v)
		}
	}
	return result
}

func main() {
	leetcode.D(func() interface{} {
		return longestPalindrome("babad")
	})
}

/**
TODO：题解用了 Manacher 算法，有些数组和变量后续优化 可达最优解

思路：
1.每扩散一次，设置中心与半径
2.每移动一次，判断是否有镜像
 */
