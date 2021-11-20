package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays

func intersection(nums1 []int, nums2 []int) []int {
	var m = map[int]struct{}{}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	var res []int
	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}
