package main

import "math/bits"

// https://leetcode.cn/problems/transform-to-chessboard

func getMoves(mask uint, cnt, n int) int {
	ones := bits.OnesCount(mask)
	if n%2 == 1 {
		if abs(n-2*ones) != 1 || abs(n-2*cnt) != 1 {
			return -1
		}
		if ones == n>>1 { // 偶数位 0
			return n/2 - bits.OnesCount(mask&0xAAAAAAAA) // 1010
		} else { // 奇数位 0
			return n/2 - bits.OnesCount(mask&0x55555555) + 1 // 0101
		}
	} else {
		if ones != n>>1 || cnt != n>>1 {
			return -1
		}
		// 偶数位 0
		cnt0 := n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		// 奇数位 0
		cnt1 := n/2 - bits.OnesCount(mask&0x55555555)
		return min(cnt0, cnt1)
	}
}

func movesToChessboard(board [][]int) int {
	bL := len(board)
	// 棋盘的第一行与第一列
	rowMask, colMask := 0, 0
	for i := 0; i < bL; i++ {
		rowMask |= board[0][i] << i
		colMask |= board[i][0] << i
	}
	rowMaskReverse := 1<<bL - 1 ^ rowMask // TODO
	colMastReverse := 1<<bL - 1 ^ colMask
	rowCnt, colCnt := 0, 0
	for i := 0; i < bL; i++ {
		rowMaskCur, colMaskCur := 0, 0
		for j := 0; j < bL; j++ {
			rowMaskCur |= board[i][j] << j
			colMaskCur |= board[j][i] << j
		}

		if rowMaskCur != rowMask && rowMaskCur != rowMaskReverse || // 检测每一行的状态是否合法
			colMaskCur != colMask && colMaskCur != colMastReverse { // 检测每一列的状态是否合法
			return -1
		}
		if rowMaskCur == rowMask {
			rowCnt++ // 记录与第一行相同的行数
		}
		if colMaskCur == colMask {
			colCnt++ // 记录与第一列相同的列数
		}
	}
	rowMoves := getMoves(uint(rowMask), rowCnt, bL)
	colMoves := getMoves(uint(colMask), colCnt, bL)
	if rowMoves == -1 || colMoves == -1 {
		return -1
	}
	return rowMoves + colMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
