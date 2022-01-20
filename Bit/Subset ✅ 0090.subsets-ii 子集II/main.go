package main

// https://leetcode-cn.com/problems/subsets-ii

// ❓ 子集II

func subsetsWithDup(nums []int) [][]int {
	var combNums [][]int
	//sort.Ints(nums)

	numsL := len(nums)
	for subset := 0; subset < 1<<numsL; subset++ {
		vis := map[int]int{}
		numsCur := []int{}
		idx := 0
		for idx < numsL {
			if subset>>idx&1 == 1 { // 是否选择
				idxPre, ok := vis[nums[idx]]
				//if 0 < idx && subset>>(idx-1)&1 == 0 && nums[idx] == nums[idx-1] {
				if ok && subset>>idxPre&1 == 0 {
					// 上一个数没有选择，本次也不选择
					break
				}
				numsCur = append(numsCur, nums[idx])
			}
			vis[nums[idx]] = idx
			idx++
		}
		if idx == numsL {
			combNums = append(combNums, numsCur)
		}
	}
	return combNums
}
