package main

// https://leetcode-cn.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}

	var res int
	for num := range m {
		if !m[num-1] {
			cur := 1
			num += 1
			for m[num] {
				num++
				cur++
			}
			if res < cur {
				res = cur
			}
		}
	}
	return res
}
