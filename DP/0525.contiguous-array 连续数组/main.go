package main

// https://leetcode-cn.com/problems/contiguous-array

func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	var counter, res int
	for i, v := range nums {
		if v == 1 {
			counter++
		} else {
			counter--
		}
		index, ok := m[counter]
		if ok {
			if res < i-index {
				res = i - index
			}
		} else {
			m[counter] = i
		}
	}
	return res
}
