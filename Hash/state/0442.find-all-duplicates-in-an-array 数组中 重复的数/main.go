package main

// https://leetcode.cn/problems/find-all-duplicates-in-an-array

// ❓ 找出 1 <= nums[i] <= numsL 数组中出现2次的元素

func findDuplicates(nums []int) []int {
	numsL := len(nums) + 1
	var numsDuplicate []int
	for _, num := range nums {
		// 索引位置
		idx := (num - 1) % numsL
		if nums[idx]/numsL == 1 {
			// 已出现过
			numsDuplicate = append(numsDuplicate, idx+1)
		}
		nums[idx] += numsL // 标记 min = n + 1
	}
	return numsDuplicate
}
