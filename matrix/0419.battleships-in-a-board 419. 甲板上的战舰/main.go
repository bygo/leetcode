package main

// https://leetcode.cn/problems/battleships-in-a-board

// ❓ 矩阵上找几条直线
// ⚠️ 1*k k*1 方式排列

// 访问后设置空
func countBattleships(board [][]byte) int {
	var cnt int
	boardL, rowL := len(board), len(board[0])
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				row[j] = '.'
				// 行计算
				for k := j + 1; k < rowL && row[k] == 'X'; k++ {
					row[k] = '.'
				}
				// 列计算
				for k := i + 1; k < boardL && board[k][j] == 'X'; k++ {
					board[k][j] = '.'
				}
				cnt++
			}
		}
	}
	return cnt
}

// 计算左上顶点
func countBattleships(board [][]byte) int {
	var cnt int
	for i, row := range board {
		for j, ch := range row {
			if ch == 'X' {
				// 假设新增
				cnt++
				if 0 < i && board[i-1][j] == 'X' || 0 < j && board[i][j-1] == 'X' {
					// 上或左 有X，证明不是左上顶点
					cnt--
				}
			}
		}
	}
	return cnt
}
