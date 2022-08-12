package main

// https://leetcode.cn/problems/two-sum/

// ❓ 两数之和

func twoSum(nums []int, target int) []int {
	var numMpIdx = map[int]int{}
	for idxRaw, numRaw := range nums {
		numHash := target - numRaw
		idxHash, ok := numMpIdx[numHash]
		if ok {
			return []int{idxHash, idxRaw}
		}
		numMpIdx[numRaw] = idxRaw
	}
	return nil
}
