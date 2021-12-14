package main

// https://leetcode-cn.com/problems/contains-duplicate

// ❓出现重复

func containsDuplicate(nums []int) bool {
	var valMpCnt = map[int]int{}
	for _, num := range nums {
		if 0 < valMpCnt[num] {
			return true
		}
		valMpCnt[num] += 1
	}
	return false
}
