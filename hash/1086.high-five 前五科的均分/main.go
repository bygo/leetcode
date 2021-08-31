package main

import "sort"

// https://leetcode-cn.com/problems/high-five

func highFive(items [][]int) [][]int {
	m := [1001][]int{}
	for _, item := range items {
		m[item[0]] = append(m[item[0]], item[1])
	}

	var res [][]int

	for i := range m {
		if m[i] != nil {
			sort.Slice(m[i], func(l, r int) bool {
				return m[i][r] < m[i][l]
			})
			var sum int
			for _, v := range m[i][:5] {
				sum += v
			}
			sum /= 5
			res = append(res, []int{i, sum})
		}
	}
	return res
}
