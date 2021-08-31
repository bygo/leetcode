package main

//
// https://leetcode-cn.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	var m = map[int]int{}
	for _, num := range nums {
		if 0 < m[num] {
			return true
		}
		m[num] += 1
	}
	return false
}
