package main

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array

func findMaximumXOR(nums []int) int {
	var res, next int
	var h = 30
	for i := h; 0 <= i; i-- {
		res <<= 1
		next = res + 1
		m := map[int]bool{}
		for _, num := range nums {
			cur := num >> i
			if m[cur] {
				res += 1
				break
			}
			m[cur^next] = true
		}
	}
	return res
}
