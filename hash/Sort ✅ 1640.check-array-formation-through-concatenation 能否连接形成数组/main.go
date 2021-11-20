package main

// https://leetcode-cn.com/problems/check-array-formation-through-concatenation

func canFormArray(arr []int, pieces [][]int) bool {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]] = i + 1
	}

	for _, piece := range pieces {
		top := len(piece) - 1
		for j := range piece {
			if m[piece[j]] == 0 || j < top && m[piece[j]]+1 != m[piece[j+1]] {
				return false
			}
		}
	}
	return true
}
