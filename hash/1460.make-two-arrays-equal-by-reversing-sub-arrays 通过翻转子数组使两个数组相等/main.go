package main

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays

func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}
	for i := range target {
		m[target[i]]++
	}

	for i := range arr {
		m[arr[i]]--
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
}
