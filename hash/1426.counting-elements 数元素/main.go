package main

// https://leetcode-cn.com/problems/counting-elements

func countElements(arr []int) int {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]]++
	}

	var res int
	for i := range m {
		if 0 < m[i+1] {
			res += m[i]
		}
	}
	return res
}
