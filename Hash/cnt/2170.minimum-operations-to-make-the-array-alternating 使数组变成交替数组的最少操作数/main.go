package main

import "sort"

// https://leetcode.cn/problems/minimum-operations-to-make-the-array-alternating

// ❓ 使数组变成交替数组的最少操作数

type pair struct {
	cnt int
	num int
}

func minimumOperations(nums []int) int {
	var numsL = len(nums)
	numMpCnt1 := map[int]int{}
	numMpCnt2 := map[int]int{}

	for idx, num := range nums {
		if idx%2 == 0 {
			numMpCnt1[num]++
		} else {
			numMpCnt2[num]++
		}
	}
	var nums1, nums2 []pair
	for num, cnt := range numMpCnt1 {
		nums1 = append(nums1, pair{
			cnt: cnt,
			num: num,
		})
	}

	for num, cnt := range numMpCnt2 {
		nums2 = append(nums2, pair{
			cnt: cnt,
			num: num,
		})
	}

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[j].cnt < nums1[i].cnt
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[j].cnt < nums2[i].cnt
	})
	for len(nums1) < 2 {
		nums1 = append(nums1, pair{})
	}
	for len(nums2) < 2 {
		nums2 = append(nums2, pair{})
	}
	if nums1[0].num != nums2[0].num {
		return numsL - nums1[0].cnt - nums2[0].cnt
	}

	return numsL - max(nums1[0].cnt+nums2[1].cnt, nums1[1].cnt+nums2[0].cnt)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
