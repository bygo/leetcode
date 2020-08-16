package main

/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

func longestValidParentheses(s string) int {
	var left, right, max int
	for _, v := range s {
		if v == 40 {
			left++
		} else {
			right++
		}
		if left == right {
			if max < left {
				max = left
			}
		} else if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 40 {
			left++
		} else {
			right++
		}
		if left == right {
			if max < right {
				max = right
			}
		} else if right < left {
			left, right = 0, 0
		}
	}
	return max * 2
}

/**
思路：左右遍历，left right类似栈抵消
*/
