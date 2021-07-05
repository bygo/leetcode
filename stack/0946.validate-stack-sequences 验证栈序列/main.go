package main

// Link: https://leetcode-cn.com/problems/validate-stack-sequences

// pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}

	c := 0 // popped 计数器
	for i := range pushed {
		stack = append(stack, pushed[i])                         // 一直进栈
		for 0 < len(stack) && stack[len(stack)-1] == popped[c] { // 相等就出栈
			c++
			stack = stack[:len(stack)-1]
		}
	}
	return c == len(pushed)
}
