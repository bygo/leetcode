package main

// https://leetcode.cn/problems/two-out-of-three

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := map[int]int{}
	var res []int

	for i := range nums1 {
		if m[nums1[i]] == 0 {
			m[nums1[i]] = 1
		}
	}

	for i := range nums2 {
		if m[nums2[i]] < 2 {
			if m[nums2[i]] == 1 {
				res = append(res, nums2[i])
			}
			m[nums2[i]] += 2
		}
	}

	for i := range nums3 {
		if m[nums3[i]] == 1 || m[nums3[i]] == 2 {
			res = append(res, nums3[i])
			m[nums3[i]] += 4
		}
	}
	return res
}
