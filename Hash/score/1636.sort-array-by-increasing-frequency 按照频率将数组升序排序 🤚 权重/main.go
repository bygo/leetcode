package main

import "sort"

// https://leetcode.cn/problems/sort-array-by-increasing-frequency

func frequencySort(nums []int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}

	for i := range nums {
		nums[i] = 201*m[nums[i]] - nums[i] + 100
	}

	sort.Ints(nums)

	for i := range nums {
		nums[i] = -(nums[i] % 201) + 100
	}
	return nums
}
