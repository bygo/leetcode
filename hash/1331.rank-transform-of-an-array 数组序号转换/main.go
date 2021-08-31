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
		_, ok := m[sorted[i]]
		if !ok {
			cnt++
		}
		m[sorted[i]] = cnt
	}

	var res []int
	for i := range arr {
		res = append(res, m[arr[i]])
	}
	return res
}
