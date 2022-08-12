package main

// https://leetcode.cn/problems/intersection-of-two-arrays-ii

// ❓ 两数组 val 交集

func intersect(nums1 []int, nums2 []int) []int {
	numMpCnt := map[int]int{}
	for _, num := range nums1 {
		numMpCnt[num]++
	}

	var numsInter []int
	for _, num := range nums2 {
		if 0 < numMpCnt[num] {
			numsInter = append(numsInter, num)
			numMpCnt[num]--
		}
	}
	return numsInter
}
