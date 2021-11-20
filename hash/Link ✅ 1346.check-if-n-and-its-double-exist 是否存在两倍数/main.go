package main

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist

func checkIfExist(arr []int) bool {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]]++
	}

	for i := range m {
		if 0 < m[i*2] {
			if i != 0 || 2 <= m[i] {
				return true
			}
		}
	}
	return false
}
