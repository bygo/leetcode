package main

import "sort"

// https://leetcode-cn.com/problems/merge-intervals

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	var cur = intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] <= cur[1] {
			if cur[1] < interval[1] {
				cur[1] = interval[1]
			}
		} else {
			res = append(res, cur)
			cur = interval
		}
	}
	res = append(res, cur)
	return res
}
