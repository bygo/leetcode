package main

// https://leetcode.cn/problems/convert-1d-array-into-2d-array

func construct2DArray(original []int, m int, n int) [][]int {
	oriL := len(original)
	if m*n != oriL {
		return nil
	}
	var groupsNums = make([][]int, m)
	for i := range groupsNums {
		groupsNums[i] = make([]int, n)
	}
	for idx, num := range original {
		idxGroup := idx / n
		idxNum := idx % n
		groupsNums[idxGroup][idxNum] = num
	}
	return groupsNums
}
