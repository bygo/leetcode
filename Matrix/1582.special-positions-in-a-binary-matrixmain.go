package Matrix

// https://leetcode.cn/problems/special-positions-in-a-binary-matrix

func numSpecial(mat [][]int) int {
	rowL := len(mat)
	colL := len(mat[0])
	rowMpCnt := make([]int, rowL)
	colMpCnt := make([]int, colL)
	for i := 0; i < rowL; i++ {
		for j := 0; j < colL; j++ {
			rowMpCnt[i] += mat[i][j]
			colMpCnt[j] += mat[i][j]
		}
	}

	var cnt int
	for i := 0; i < rowL; i++ {
		for j := 0; j < colL; j++ {
			if mat[i][j] != 1 {
				continue
			}
			if rowMpCnt[i] != 1 {
				continue
			}
			if colMpCnt[j] != 1 {
				continue
			}
			cnt++
		}
	}
	return cnt
}
