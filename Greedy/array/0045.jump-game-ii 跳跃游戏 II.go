package array

// https://leetcode.cn/problems/jump-game-ii

func jump(nums []int) int {
	var steps, idxCur, idxFar int
	top := len(nums) - 1
	for idx := 0; idx < top; idx++ {
		idxFar = max(idxFar, idx+nums[idx])
		if idxCur == idx {
			idxCur = idxFar
			steps++
		}
	}
	return steps
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func jump(nums []int) int {
	var steps int
	idxCur := len(nums) - 1
	for 0 < idxCur {
		for idx := 0; idx < idxCur; idx++ {
			if idxCur <= idx+nums[idx] {
				idxCur = idx
				steps++
				break
			}
		}
	}
	return steps
}

func jump(nums []int) int {
	var steps int
	idxCur := len(nums) - 1
	for 0 < idxCur {
		for idx := 0; idx < idxCur; idx++ {
			if idxCur <= idx+nums[idx] {
				idxCur = idx
				steps++
				break
			}
		}
	}
	return steps
}
