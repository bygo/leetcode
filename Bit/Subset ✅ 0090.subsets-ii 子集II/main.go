package main

// https://leetcode-cn.com/problems/subsets-ii

// ❓ 子集II

func subsetsWithDup(nums []int) [][]int {
	var combNums [][]int
	//sort.Ints(nums)

	numsL := len(nums)
	subsetMax := 1 << numsL
	for subset := 0; subset < subsetMax; subset++ {
		numMpIdx := map[int]int{}
		numsCur := []int{}
		idx := 0
		for idx < numsL {
			if subset>>idx&1 == 1 { // 是否选择
				idxPre, ok := numMpIdx[nums[idx]]
				//if 0 < idx && subset>>(idx-1)&1 == 0 && nums[idx] == nums[idx-1] {
				if ok && subset>>idxPre&1 == 0 {
					// 上一个数没有选择，本次也不选择
					break
				}
				numsCur = append(numsCur, nums[idx])
			}
			numMpIdx[nums[idx]] = idx
			idx++
		}
		if idx == numsL {
			combNums = append(combNums, numsCur)
		}
	}
	return combNums
}

func subsets(nums []int) [][]int {
	var combNums [][]int
	numsL := len(nums)
	subsetMax := 1 << numsL
	for subset := 0; subset < subsetMax; subset++ {
		numMpIdx := map[int]int{}
		numsCur := []int{}
		idx := 0
		for idx < numsL {
			if subset>>idx&1 == 1 {
				idxPre, ok := numMpIdx[nums[idx]]
				if ok && subset>>idxPre&1 == 0 {
					break
				}
				numsCur = append(numsCur, nums[idx])
			}
		}
		if idx == numsL {
			combNums = append(combNums, numsCur)
		}
	}
	return combNums
}
