package main

// https://leetcode-cn.com/problems/search-a-2d-matrix

func searchMatrix(matrix [][]int, target int) bool {
	l1 := len(matrix)
	l2 := len(matrix[0])
	l, r := 0, l1*l2
	for l < r {
		mid := l + (r-l)/2
		num := matrix[mid/l2][mid%l2]
		if num < target {
			l = mid + 1
		} else if target < num {
			r = mid
		} else if target == num {
			return true
		}
	}
	return false
}
