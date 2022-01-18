package main

// https://leetcode-cn.com/problems/combinations

func combine(n int, needL int) [][]int {
	var combsNums [][]int
	var nums []int
	var dfs func(num int)
	dfs = func(num int) {
		numsL := len(nums)
		if numsL == needL {
			combsNums = append(combsNums, append([]int{}, nums...))
			return
		}

		if n-num < needL-numsL {
			return
		}

		nums = append(nums, num+1)
		dfs(num + 1)
		nums = nums[:len(nums)-1]

		// 跳过当前
		dfs(num + 1)
	}
	dfs(0)
	return combsNums
}
