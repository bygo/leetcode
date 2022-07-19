package main

// https://leetcode.cn/problems/contiguous-array

func findMaxLength(nums []int) int {
	cntMpIdx := map[int]int{0: -1}
	var cnt, distMax int
	for i, num := range nums {
		if num == 1 {
			cnt++
		} else {
			cnt--
		}
		// 前缀溢出1的个数

		idx, ok := cntMpIdx[cnt]
		if ok {
			distCur := i - idx
			if distMax < distCur {
				distMax = distCur
			}
		} else {
			cntMpIdx[cnt] = i
		}
	}
	return distMax
}
