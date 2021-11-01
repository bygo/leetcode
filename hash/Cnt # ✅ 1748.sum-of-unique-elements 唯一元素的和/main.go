package main

// https://leetcode-cn.com/problems/sum-of-unique-elements

func sumOfUnique(nums []int) int {
	var m = map[int]int{}
	var res int
	for _, v := range nums {
		if m[v] == 0 {
			res += v
			m[v] = 1
		} else if m[v] == 1 {
			res -= v
			m[v] = 2
		}
	}
	return res
}
