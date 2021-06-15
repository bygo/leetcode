package main

//Link: https://leetcode-cn.com/problems/intersection-of-two-arrays

func intersection(nums1 []int, nums2 []int) []int {
	var m = map[int]int{}
	for _, v := range nums1 {
		m[v] = 1
	}

	var res []int
	for _, v := range nums2 {
		if m[v] == 1 {
			res = append(res, v)
			m[v] = 0
		}
	}
	return res
}
