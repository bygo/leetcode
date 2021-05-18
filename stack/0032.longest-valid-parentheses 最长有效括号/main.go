package main

//Link:https://leetcode-cn.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) int {
	var left, right, max int
	for _, v := range s {
		if v == '(' {
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
	for i := len(s) - 1; 0 < i; i-- {
		if s[i] == '(' {
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

func longestValidParentheses(s string) int {
	res := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 刚好全部抵消
				stack = append(stack, i) // 重新记录前置下标
			} else {
				res = max(res, i-stack[len(stack)-1]) //  每次抵消 计算最长长度
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
思路：左右遍历，left right类似栈抵消
*/
