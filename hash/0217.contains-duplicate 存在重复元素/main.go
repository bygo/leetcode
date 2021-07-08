package main

// Link: https://leetcode-cn.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	var set = map[int]int{}
	for _, num := range nums {
		set[num] += 1
	}
	for _, cnt := range set {
		if 1 < cnt {
			return true
		}
	}
	return false
}
