package main

import "sort"

// https://leetcode.cn/problems/assign-cookies

func findContentChildren(greed, size []int) int {
	sort.Ints(greed)
	sort.Ints(size)
	var cnt int
	gL, sL := len(greed), len(size)
	for idxG, idxS := 0, 0; idxG < gL && idxS < sL; idxG++ {
		for idxS < sL && size[idxS] < greed[idxG] {
			idxS++ // TODO 贪心
		}
		if idxS < sL { // TODO 分配
			cnt++
			idxS++
		}
	}
	return cnt
}
