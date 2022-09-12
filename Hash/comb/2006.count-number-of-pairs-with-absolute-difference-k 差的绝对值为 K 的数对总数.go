package main

// https://leetcode.cn/problems/count-number-of-pairs-with-absolute-difference-k

// ❓ 差为K的数对总数

func countKDifference(nums []int, k int) int {
	numMpCnt := map[int]int{}
	var cntDiff int
	for _, num := range nums {
		cntDiff += numMpCnt[num-k] + numMpCnt[num+k]
		numMpCnt[num]++
	}
	return cntDiff
}
