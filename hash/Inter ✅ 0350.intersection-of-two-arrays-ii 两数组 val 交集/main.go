package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

//
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for i := range nums1 {
		m[nums1[i]]++
	}

	var res []int
	for i := range nums2 {
		if 0 < m[nums2[i]] {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}
