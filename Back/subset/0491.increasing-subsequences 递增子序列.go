package main

// https://leetcode.cn/problems/increasing-subsequences

// ❓ 递增子序列

func findSubsequences(nums []int) [][]int {
	var combNums [][]int
	var numsCur []int
	var dfs func(idx, numPre int)
	var numsL = len(nums)
	dfs = func(idx, numPre int) {
		if numsL == idx {
			if 2 <= len(numsCur) {
				combNums = append(combNums, append([]int{}, numsCur...))
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
		// 跳过相等
		if numPre != nums[idx] {
			// 前置跳过，后置 numPre <= nums[idx] 拼接
			dfs(idx+1, numPre)
		}
	}
	dfs(0, -101)
	return combNums
}
