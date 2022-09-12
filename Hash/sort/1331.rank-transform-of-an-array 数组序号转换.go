package main

import "sort"

// https://leetcode.cn/problems/rank-transform-of-an-array

// ❓ 数组序号转换

func arrayRankTransform(arr []int) []int {
	arrL := len(arr)
	arrSorted := make([]int, arrL)
	copy(arrSorted, arr)
	sort.Ints(arrSorted)

	numMpCnt := map[int]int{}
	cnt := 0
	// 排名
	for _, num := range arrSorted {
		if numMpCnt[num] == 0 {
			cnt++
		}
		numMpCnt[num] = cnt
	}

	// 设置序号
	for i := range arr {
		arr[i] = numMpCnt[arr[i]]
	}
	return arr
}
