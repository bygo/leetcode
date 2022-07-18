package main

// https://leetcode.cn/problems/largest-number-at-least-twice-of-others

// ❓ 至少是其他数字两倍的最大数

func dominantIndex(nums []int) int {
	var max1, max2, idx int
	for i, num := range nums {
		if max1 < num {
			max2 = max1
			max1 = num
			idx = i
		} else if max2 < num {
			max2 = num
		}
	}

	if max2*2 <= max1 {
		return idx
	}
	return -1
}
