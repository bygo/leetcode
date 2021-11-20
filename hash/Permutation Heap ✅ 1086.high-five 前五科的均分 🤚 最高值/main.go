package main

import "sort"

// https://leetcode-cn.com/problems/high-five

func highFive(items [][]int) [][]int {
	m1 := map[int]int{}
	m2 := map[int]int{}
	slice := [][]int{}
	getSlice := func(id int) *[]int {
		_, ok := m1[id]
		if !ok {
			l := len(slice)
			m1[id] = l
			m2[l] = id
			slice = append(slice, []int{})
		}
		return &slice[m1[id]]
	}
	getId := func(index int) int {
		return m2[index]
	}
	for _, item := range items {
		s := getSlice(item[0])
		*s = append(*s, item[1])
	}

	var res [][]int
	for i := range slice {
		sort.Slice(slice[i], func(l, r int) bool {
			return slice[i][r] < slice[i][l]
		})
		var sum int
		for _, v := range slice[i][:5] {
			sum += v
		}
		sum /= 5
		res = append(res, []int{getId(i), sum})
	}

	sort.Slice(res, func(l, r int) bool {
		return res[l][0] < res[r][0]
	})
	return res
}
