package main

// https://leetcode.cn/problems/check-array-formation-through-concatenation

// ❓ 能否连接成数组

func canFormArray(arr []int, pieces [][]int) bool {
	numMpIdx := map[int]int{}
	for i, num := range arr {
		numMpIdx[num] = i
	}

	for _, piece := range pieces {
		top := len(piece) - 1
		for i := range piece {
			numCur := piece[i]
			_, ok := numMpIdx[numCur]
			if !ok {
				// 不存在
				return false
			}
			if i < top {
				// 当前值的索引+1 是否等于下一个值的索引
				numNext := piece[i+1]
				if numMpIdx[numCur]+1 != numMpIdx[numNext] {
					// 无法连接
					return false
				}
			}
		}
	}
	return true
}
