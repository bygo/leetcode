package main

// https://leetcode.cn/problems/combinations

func combine(numMax int, needL int) [][]int {
	var combNums [][]int
	var nums []int
	var dfs func(num int)
	dfs = func(num int) {
		numsL := len(nums)
		if numsL == needL {
			combNums = append(combNums, append([]int{}, nums...))
			return
		}

		if numMax-num-1 < needL-numsL {
			return
		}
		nums = append(nums, num)
		dfs(num + 1)
		nums = nums[:len(nums)-1]
		dfs(num + 1)
	}
	dfs(1)
	return combNums
}
