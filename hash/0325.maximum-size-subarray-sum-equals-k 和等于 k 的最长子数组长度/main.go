package main

// https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k

func maxSubArrayLen(nums []int, k int) int {
	m := map[int]int{0: -1}
	var pre, max, cur int
	for i := range nums {
		pre += nums[i]
		index, ok := m[pre-k]
		if ok {
			cur = i - index
			if max < cur {
				max = cur
			}
		}

		_, ok = m[pre]
		if !ok {
			m[pre] = i
		}
	}
	return max
}
