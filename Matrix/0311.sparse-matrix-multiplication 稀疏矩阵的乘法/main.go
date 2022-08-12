package main

// https://leetcode.cn/problems/sparse-matrix-multiplication

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	l1 := len(mat2)
	l2 := len(mat2[0])

	var res = make([][]int, 0, len(mat1))
	for i := range mat1 {
		for j := 0; j < l2; j++ {
			var cur int
			for k := 0; k < l1; k++ {
				cur += mat1[i][k] * mat2[k][j] // 行列相乘 后相加
			}
			res[i] = append(res[i], cur)
		}
	}
	return res
}
