package main

import (
	"math"
)

// https://leetcode-cn.com/problems/increasing-subsequences

func findSubsequences(nums []int) [][]int {
	var combNums [][]int
	numsL := len(nums)
	subsetMax := 1 << numsL
	hashMp := map[uint]struct{}{}
	for num := 0; num < subsetMax; num++ {
		numsCur := []int{}
		idx := 0
		numPre := -101
		subset := num
		for idx < numsL {
			if subset>>idx&1 == 1 && numPre <= nums[idx] { // 是否选择 且合法
				numsCur = append(numsCur, nums[idx])
				numPre = nums[idx]
			}
			idx++
		}

		if idx == numsL && 2 <= len(numsCur) {
			h := hash(numsCur)
			_, ok := hashMp[h]
			if !ok {
				hashMp[h] = struct{}{}
				combNums = append(combNums, numsCur)
			}
		}
	}
	return combNums
}

func hash(nums []int) uint {
	var h uint
	for _, num := range nums {
		h = h*131 + uint(num+101)
	}
	return h
}

func findSubsequences_(nums []int) [][]int {
	var subsetNums [][]int
	var numsCur []int
	var dfs func(idx, numPre int)
	var numsL = len(nums)
	dfs = func(idx, numPre int) {
		if numsL == idx {
			numsCurL := len(numsCur)
			if 2 <= numsCurL {
				subsetNums = append(subsetNums, append([]int{}, numsCur...))
			}
			return
		}

		// 5,(7),7,7
		// 5,(7),7
		// 5,(7)
		// 7,7,7
		// 7,7

		// <= 形成分支
		if numPre <= nums[idx] {
			numsCur = append(numsCur, nums[idx])
			dfs(idx+1, nums[idx])
			numsCur = numsCur[:len(numsCur)-1]
		}
		// 跳过当前
		if numPre != nums[idx] {
			// 前置跳过，后置 numPre <= nums[idx] 拼接
			dfs(idx+1, numPre)
		}
	}
	dfs(0, math.MinInt32)
	return subsetNums
}
