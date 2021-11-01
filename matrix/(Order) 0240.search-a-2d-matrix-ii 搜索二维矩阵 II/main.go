package main

import "sort"

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii

func searchMatrix(matrix [][]int, target int) bool {
	for _, i := range matrix {
		j := sort.SearchInts(i, target)
		if j < len(i) && i[j] == target {
			return true
		}
	}
	return false
}

// 舍弃流
func searchMatrix(matrix [][]int, target int) bool {
	l1, l2 := len(matrix), len(matrix[0])
	i, j := 0, l2-1
	for i < l1 && 0 <= j {
		cur := matrix[i][j]
		if cur < target {
			i++
		} else if target < cur {
			j--
		} else if target == cur {
			return true
		}
	}
	return false
}
