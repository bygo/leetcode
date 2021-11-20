package main

// https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k

func maxSubArrayLen(nums []int, k int) int {
	m := map[int]int{0: -1}
	var pre, max, cur int
	for r := range nums {
		pre += nums[r]
		l, ok := m[pre-k]
		if ok {
			cur = r - l
			if max < cur {
				max = cur
			}
		}

		_, ok = m[pre]
		if !ok {
			m[pre] = r
		}
	}
	return max
}
