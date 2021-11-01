package main

// https://leetcode-cn.com/problems/number-of-good-pairs

func numIdenticalPairs(nums []int) int {
	var m = map[int]int{}
	var res int
	for _, v := range nums {
		if 0 < m[v] {
			res += m[v]
		}
		m[v]++
	}
	return res
}
