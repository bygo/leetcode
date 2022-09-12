package main

// https://leetcode.cn/problems/increasing-subsequences

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

		// <= 合法分支
		if numPre <= nums[idx] {
			numsCur = append(numsCur, nums[idx])
			dfs(idx+1, nums[idx])
			numsCur = numsCur[:len(numsCur)-1]
		}
		// 去重
		if numPre == nums[idx] {
			return
		}
		// 跳过当前
		dfs(idx+1, numPre)
	}
	dfs(0, -101)
	return combNums
}
