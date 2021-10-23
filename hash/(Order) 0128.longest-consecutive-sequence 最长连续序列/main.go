package main

// https://leetcode-cn.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}

	var cur, res int
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
