package main

// Link: https://leetcode-cn.com/problems/sum-of-unique-elements

func sumOfUnique(nums []int) int {
	var m = map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}

	var res int
	for k, v := range m {
		if v == 1 {
			res += k
		}
	}
	return res
}
