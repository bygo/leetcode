package main

import "sort"

// https://leetcode-cn.com/problems/find-original-array-from-doubled-array

func findOriginalArray(changed []int) []int {
	m := map[int]int{}
	sort.Ints(changed)
	var res []int
	for _, num := range changed {
		if m[num] == 0 {
			m[num*2]++
			res = append(res, num)
		} else {
			m[num]--
		}
	}
	for _, num := range m {
		if num != 0 {
			return nil
		}
	}
	return res
}
