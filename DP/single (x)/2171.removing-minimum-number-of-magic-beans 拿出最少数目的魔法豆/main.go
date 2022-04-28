package main

import "sort"

// https://leetcode-cn.com/problems/removing-minimum-number-of-magic-beans

// 📚
func minimumRemoval(beans []int) int64 {
	beansL := len(beans)
	sort.Ints(beans)
	var preSum, numMax int
	for idx, num := range beans {
		preSum += num
		numCur := (beansL - idx) * num // 剩余豆子，idx~beansL 留下，
		if numMax < numCur {
			numMax = numCur
		}
	}
	return int64(preSum - numMax)
}
