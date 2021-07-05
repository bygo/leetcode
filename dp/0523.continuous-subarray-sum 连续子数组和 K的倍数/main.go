package main

// Link: https://leetcode-cn.com/problems/continuous-subarray-sum

// 前缀余数
func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	m := map[int]int{0: -1}
	r := 0
	for i, v := range nums {
		r = (r + v) % k
		index, ok := m[r]
		if ok {
			if 2 <= i-index {
				return true
			}
		} else {
			m[r] = i
		}
	}
	return false
}
