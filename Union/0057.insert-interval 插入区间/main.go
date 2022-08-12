package main

import "sort"

// https://leetcode.cn/problems/insert-interval

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	start, end := newInterval[0], newInterval[1]
	var l, r = 0, n
	for l < r {
		mid := (l + r) >> 1
		cur := intervals[mid]
		if cur[1] < start { //找最左边界
			l = mid + 1
		} else if start <= cur[1] {
			r = mid
		}
	}

	// 结束时在右侧
	left := l
	left = sort.Search(n, func(i int) bool {
		return start <= intervals[i][1]
	})
	r = n
	for l < r {
		mid := (l + r) >> 1
		cur := intervals[mid]
		if end < cur[0] { // 找最右边界
			r = mid
		} else if cur[0] <= end {
			l = mid + 1
		}
	}
	// 结束时在右侧，需 - 1
	right := r - 1
	right = sort.Search(n, func(i int) bool {
		return end < intervals[i][0]
	}) - 1
	if left <= right { // 小于不用判断
		start = min(start, intervals[left][0])
		end = max(end, intervals[right][1])
	}
	var res [][]int
	res = append(res, intervals[:left]...)

	res = append(res, []int{start, end})

	res = append(res, intervals[right+1:]...)

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
