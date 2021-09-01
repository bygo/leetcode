package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

func findLucky(arr []int) int {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]]++
	}

	var max = -1
	for i := range m {
		if m[i] == i && max < i {
			max = i
		}
	}
	return max
}
