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
	m := []rune{'^', '#'}
	for _, v := range s {
		m = append(m, v, '#')
	}
	m = append(m, '$')
	p := make([]int, len(m))
	var maxCenter int
	var maxRight int
	var center int
	var right int
	for k, _ := range m {
		if k == 0 || k == len(m)-1 {
			continue
		}

		if maxRight >= len(m)-k { //不够长，break
			break
		}

		if k < right { //镜像
			mirror := p[center*2-k]
			if mirror+k > right {
				p[k] = right - k
			} else {
				p[k] = p[center*2-k]
			}
		} else {
			p[k] = 1
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

	result := make([]rune,0)
	for _, v := range m[maxCenter-maxRight+1 : maxCenter+maxRight] {
		if v != '#' {
			result = append(m, v)
		}
	}
	return string(result)
}

func main() {
	leetcode.D(func() interface{} {
		return longestPalindrome("bbasdfasdfabbabasdfasdfasbasdfafasdfaskasdfjkasjdfaksdjfkasjdfkajsdkfjaskdjfadsf")
	})
}

/**
思路：
1.每扩散一次，设置中心与半径
2.每移动一次，判断是否有镜像
 */
