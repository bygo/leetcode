package main

// 相对交集 val 取唯一值
// https://leetcode-cn.com/problems/intersection-of-two-arrays

func intersection(nums1 []int, nums2 []int) []int {
	var m = map[int]struct{}{}
	for _, num := range nums1 {
		m[num] = struct{}{}
	}

	var res []int
	for _, num := range nums2 {
		_, ok := m[num]

		if ok {
			res = append(res, num)
			delete(m, num)
		}
	}
	return res
}
