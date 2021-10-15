package main

import "sort"

// https://leetcode-cn.com/problems/find-original-array-from-doubled-array

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	m := map[int]int{}
	var res []int
	for i := range changed {
		if m[changed[i]] == 0 {
			m[changed[i]*2]++
			res = append(res, changed[i])
		} else {
			m[changed[i]]--
		}
	}

	for i := range m {
		if m[i] != 0 {
			return nil
		}
	}

	return res
}
