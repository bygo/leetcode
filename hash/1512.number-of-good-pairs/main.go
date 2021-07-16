package main

//Title: Number of Good Pairs
// https://leetcode-cn.com/problems/number-of-good-pairs

func numIdenticalPairs(nums []int) int {
	var m = map[int]int{}
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	var res int
	for _, v := range m {
		res += v * (v - 1) / 2
	}
	return res
}
