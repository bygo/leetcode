package main

// https://leetcode-cn.com/problems/sparse-matrix-multiplication

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	var res [][]int
	l1 := len(mat2)
	l2 := len(mat2[0])
	for i := range mat1 {
		res = append(res, []int{})
		for j := 0; j < l2; j++ {
			var cur int
			for k := 0; k < l1; k++ {
				cur += mat1[i][k] * mat2[k][j]
			}
			res[i] = append(res[i], cur)
		}
	}
	return res
}
