package main

// https://leetcode-cn.com/problems/longest-harmonious-subsequence

func findLHS(nums []int) int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}
	var res int
	for num, v := range m {
		if 0 < m[num+1] {
			cur := m[num+1] + v
			if res < cur {
				res = cur
			}
		}
	}
	return res
}

// 一次扫描
func findLHS(nums []int) int {
	m := map[int]int{}
	var res int
	for _, num := range nums {
		m[num]++
		if 0 < m[num+1] {
			cur := m[num+1] + m[num]
			if res < cur {
				res = cur
			}
		}
		if 0 < m[num-1] {
			cur := m[num-1] + m[num]
			if res < cur {
				res = cur
			}
		}
	}
	return res
}
