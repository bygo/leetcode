package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for i, num := range nums {
		if k, ok := set[target-num]; ok {
			return []int{i, k}
		}
		set[num] = i
	}
	return nil
}
