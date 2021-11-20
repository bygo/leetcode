package main

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays

func canBeEqual(target []int, arr []int) bool {
	l1 := len(target)
	l2 := len(arr)
	if l1 != l2 {
		return false
	}
	m := map[int]int{}
	for i := 0; i < l1; i++ {
		m[target[i]]++
		m[arr[i]]--
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
}
