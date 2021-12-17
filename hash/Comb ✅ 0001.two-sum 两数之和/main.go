package main

// https://leetcode-cn.com/problems/two-sum/

// ❓ 两数之和

func twoSum(nums []int, target int) []int {
	var numMpIdx = map[int]int{}
	for idxRaw, valRaw := range nums {
		idxHash, ok := numMpIdx[target-valRaw]
		if ok {
			return []int{idxHash, idxRaw}
		}
		numMpIdx[valRaw] = idxRaw
	}
	return nil
}
