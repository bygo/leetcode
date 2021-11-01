package main

import "sort"

// https://leetcode-cn.com/problems/rank-transform-of-an-array

func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	m := map[int]int{}
	cnt := 0
	for i := range sorted {
		if m[sorted[i]] == 0 {
			cnt++
		}
		m[sorted[i]] = cnt
	}

	for i := range arr {
		arr[i] = m[arr[i]]
	}
	return arr
}
