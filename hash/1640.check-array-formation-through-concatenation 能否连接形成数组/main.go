package main

// https://leetcode-cn.com/problems/check-array-formation-through-concatenation

func canFormArray(arr []int, pieces [][]int) bool {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]] = i + 1
	}

	for _, piece := range pieces {
		for i := range piece {
			if m[piece[i]] == 0 {
				return false
			}
			if i < len(piece)-1 {
				if m[piece[i]]+1 != m[piece[i+1]] {
					return false
				}
			}
		}
	}
	return true
}
