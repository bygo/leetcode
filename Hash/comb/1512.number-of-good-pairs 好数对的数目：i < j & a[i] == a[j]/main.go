package main

// https://leetcode.cn/problems/number-of-good-pairs

// ❓ 好数对的数目
// ⚠️ i < j && nums[i] == nums[j]

func numIdenticalPairs(nums []int) int {
	var numMpCnt = map[int]int{}
	var cntGood int
	for _, num := range nums {
		cntGood += numMpCnt[num]
		numMpCnt[num]++
	}
	return cntGood
}
