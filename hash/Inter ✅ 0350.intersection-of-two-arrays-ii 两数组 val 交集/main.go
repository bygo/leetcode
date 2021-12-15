package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

// ❓ 两数组 val 交集

func intersect(nums1 []int, nums2 []int) []int {
	numMpCnt := map[int]int{}
	for i := range nums1 {
		numMpCnt[nums1[i]]++
	}

	var res []int
	for _, num := range nums2 {
		if 0 < numMpCnt[num] {
			res = append(res, num)
			numMpCnt[num]--
		}
	}
	return res
}
