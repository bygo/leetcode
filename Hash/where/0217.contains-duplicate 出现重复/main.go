package main

// https://leetcode.cn/problems/contains-duplicate

// ❓ 出现重复

func containsDuplicate(nums []int) bool {
	var numMpCnt = map[int]int{}
	for _, num := range nums {
		if 0 < numMpCnt[num] {
			return true
		}
		numMpCnt[num] += 1
	}
	return false
}
