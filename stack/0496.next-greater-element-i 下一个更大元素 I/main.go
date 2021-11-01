package main

// https://leetcode-cn.com/problems/next-greater-element-i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var m = map[int]int{}
	var stack = []int{}

	for _, v := range nums2 {
		top := len(stack) - 1
		for -1 < top && stack[top] < v {
			m[stack[top]] = v // 下一个更大的值
			stack = stack[:top]
			top--
		}
		stack = append(stack, v)
	}

	var res []int
	for _, v := range nums1 {
		index, ok := m[v]
		if ok {
			res = append(res, index)
		} else {
			res = append(res, -1)
		}
	}
	return res
}
