package main

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

// ❓ 等值数组中 消失的数

func findDisappearedNumbers(nums []int) []int {
	var numsDis []int
	numsL := len(nums)
	for _, num := range nums {
		// 索引标记
		idx := (num - 1) % numsL
		nums[idx] += numsL
	}

	for num, numState := range nums {
		if numState <= numsL {
			// 代表数字不存在
			numsDis = append(numsDis, num+1)
		}
	}
	return numsDis
}
