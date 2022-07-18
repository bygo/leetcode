package main

// https://leetcode.cn/problems/next-greater-element-ii

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nums = append(nums, nums[:n-1]...)
	var stack []int
	for i, v := range nums {
		top := len(stack) - 1
		for -1 < top && nums[stack[top]] < v {
			nums[stack[top]] = v // 出栈后 虽然值变动了 不过也不管了 不影响
			stack = stack[:top]
			top--
		}
		stack = append(stack, i)
	}

	for _, index := range stack {
		if n <= index {
			break
		}
		nums[index] = -1
	}
	return nums[:n]
}
