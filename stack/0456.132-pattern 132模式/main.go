package main

// Link: https://leetcode-cn.com/problems/132-pattern

// 3,1,4,2
// i j k   i最大 k次大
func find132pattern(nums []int) bool {
	var stack []int
	var k = -1 << 63
	var i = len(nums) - 1
	for -1 < i {
		if nums[i] < k { // 小于上次 pop的值
			return true
		}
		top := len(stack) - 1
		for -1 < top && stack[top] < nums[i] { // 递减栈 栈顶 = j   上一次 pop = k
			k = stack[top]
			stack = stack[:top]
			top--
		}
		stack = append(stack, nums[i]) // 栈顶j 必比 k 大
		i--
	}
	return false
}
