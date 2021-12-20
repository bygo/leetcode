package main

// https://leetcode-cn.com/problems/remove-k-digits

func removeKdigits(num string, k int) string {
	stack := []rune{}
	for _, v := range num {
		for 0 < k && 0 < len(stack) && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, v)
	}
	stack = stack[:len(stack)-k]
	for 0 < len(stack) && stack[0] == '0' {
		stack = stack[1:]
	}

	if 0 == len(stack) {
		return "0"
	}

	return string(stack)
}
