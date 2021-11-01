package main

// https://leetcode-cn.com/problems/0H97ZC/

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := [1001]int{}
	for i := range arr1 {
		m[arr1[i]]++
	}

	var res []int
	for i := range arr2 {
		for 0 < m[arr2[i]] {
			res = append(res, arr2[i])
			m[arr2[i]]--
		}
	}

	for i := range m {
		for 0 < m[i] {
			res = append(res, i)
			m[i]--
		}
	}
	return res
}
