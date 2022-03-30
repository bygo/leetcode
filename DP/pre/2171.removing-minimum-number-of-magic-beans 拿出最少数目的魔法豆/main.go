package main

import "sort"

// https://leetcode-cn.com/problems/removing-minimum-number-of-magic-beans

// ❓ 拿出最少数目的魔法豆
func minimumRemoval(beans []int) int64 {
	beansL := len(beans)
	sort.Ints(beans)
	var sum, numMax int
	for idx, num := range beans {
		sum += num
		numCur := (beansL - idx) * num
		if numMax < numCur {
			numMax = numCur
		}
	}
	return int64(sum - numMax)
}
