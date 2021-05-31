package main

//Link:https://leetcode-cn.com/problems/longest-valid-parentheses/

// 左右计数器,
func longestValidParentheses(s string) int {
	var l = len(s)
	var max int

	left, right := 0, 0
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if max < left {
				max = left
			}
		} else if left < right { // 不合法 重新计算
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := l - 1; 0 < i; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if max < right {
				max = right
			}
		} else if right < left { // 不合法 重新计算
			left, right = 0, 0
		}
	}
	return max * 2
}

// 索引入栈，抵消计算
func longestValidParentheses(s string) int {
	var res int
	stack := []int{-1} // -1 前置索引

	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' { //
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 刚好全部抵消
				stack = append(stack, i) // 重新记录前置索引
			} else {
				res = max(res, i-stack[len(stack)-1]) //  每次抵消 计算最长长度
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/**
思路：左右遍历，left right类似栈抵消
*/
