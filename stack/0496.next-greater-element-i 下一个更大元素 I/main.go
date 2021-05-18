package main

//Link: https://leetcode-cn.com/problems/next-greater-element-i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	var m = map[int]int{}
	for _, v := range nums2 {
		top := len(stack) - 1
		for 0 <= top && stack[top] < v { //单调递减
			m[stack[top]] = v
			stack = stack[:top]
			top--
		}
		stack = append(stack, v)
	}

	var res []int
	for _, v := range nums1 {
		v, ok := m[v]
		if ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
