package main

// https://leetcode-cn.com/problems/contains-duplicate-ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = map[int]int{}
	for r, num := range nums {
		l, ok := m[num]
		if ok && r-l <= k {
			return true
		}
		m[num] = r
	}
	return false
}
