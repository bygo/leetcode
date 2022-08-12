package main

// https://leetcode.cn/problems/contains-duplicate-ii

// ❓ 是否存在两个重复元素距离小于等于K

func containsNearbyDuplicate(nums []int, k int) bool {
	var numMpIdx = map[int]int{}
	for idxCur, num := range nums {
		idxPre, ok := numMpIdx[num]
		if ok && idxCur-idxPre <= k {
			return true
		}
		numMpIdx[num] = idxCur
	}
	return false
}
