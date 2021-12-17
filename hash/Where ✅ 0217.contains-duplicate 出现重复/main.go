package main

// https://leetcode-cn.com/problems/contains-duplicate

// ❓ 出现重复

func containsDuplicate(nums []int) bool {
	var valMpCnt = map[int]int{}
	for _, val := range nums {
		if 0 < valMpCnt[val] {
			return true
		}
		valMpCnt[val] += 1
	}
	return false
}
