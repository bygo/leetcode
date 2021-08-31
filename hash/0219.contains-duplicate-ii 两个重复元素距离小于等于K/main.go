package main

// https://leetcode-cn.com/problems/contains-duplicate-ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = map[int]int{}
	for i, num := range nums {
		v, ok := m[num]
		if ok && i-v <= k {
			return true
		}
		m[num] = i
	}
	return false
}
