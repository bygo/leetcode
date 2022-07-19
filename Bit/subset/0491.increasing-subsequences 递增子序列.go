package main

// https://leetcode.cn/problems/increasing-subsequences

// ❓ 递增子序列

func findSubsequences(nums []int) [][]int {
	var combNums [][]int
	numsL := len(nums)
	subsetMax := 1 << numsL
	hashMp := map[uint]bool{}
	for subset := 0; subset < subsetMax; subset++ {
		numsCur := []int{}
		numPre := -101
		for idx := 0; idx < numsL; idx++ {
			if subset>>idx&1 == 0 || nums[idx] < numPre {
				continue
			}

			numsCur = append(numsCur, nums[idx])
			numPre = nums[idx]
		}
		if len(numsCur) < 2 {
			continue
		}

		h := hash(numsCur)
		if hashMp[h] {
			continue
		}
		hashMp[h] = true
		combNums = append(combNums, numsCur)
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
