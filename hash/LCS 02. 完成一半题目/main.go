package main

import "sort"

func halfQuestions(questions []int) int {
	m := make([]int, 1001)
	half := len(questions) / 2
	for _, question := range questions {
		m[question]++
	}
	sort.Slice(m, func(i, j int) bool {
		return m[j] < m[i]
	})
	var res int
	for i := range m {
		if half <= 0 {
			return res
		}
		half -= m[i]
		res++
	}
	return res
}
