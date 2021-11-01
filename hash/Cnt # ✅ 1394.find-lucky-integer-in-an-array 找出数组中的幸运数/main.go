package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

func findLucky(arr []int) int {
	m := [501]int{}
	for i := range arr {
		m[arr[i]]++
	}

	for i := 500; 0 < i; i-- {
		if m[i] == i {
			return i
		}
	}
	return -1
}
