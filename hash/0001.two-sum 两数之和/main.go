package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for rawK, rawV := range nums {
		hashK, ok := m[target-rawV]
		if ok {
			return []int{hashK, rawK}
		}
		m[rawV] = rawK
	}
	return nil
}
