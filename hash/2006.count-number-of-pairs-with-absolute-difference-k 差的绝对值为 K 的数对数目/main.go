package main

// https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k

func countKDifference(nums []int, k int) int {
	m := map[int]int{}
	var res int
	for _, num := range nums {
		res += m[num-k] + m[num+k]
		m[num]++
	}
	return res
}
