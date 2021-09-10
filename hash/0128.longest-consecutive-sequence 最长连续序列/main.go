package main

// https://leetcode-cn.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for i := range nums {
		m[nums[i]] = true
	}

	var res, cur int
	for num := range m {
		if !m[num-1] {
			cur = 0
			for m[num] {
				cur++
				num++
			}
			if res < cur {
				res = cur
			}

		}
	}
	return res
}
