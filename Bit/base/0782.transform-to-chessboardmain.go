package main

import "math/bits"

// https://leetcode.cn/problems/transform-to-chessboard

func getMoves(mask uint, cnt, n int) int {
	ones := bits.OnesCount(mask)
	half := n / 2
	if n%2 == 1 {
		if abs(n-2*ones) != 1 || abs(n-2*cnt) != 1 {
			return -1
		}
		if ones == half {
			// 1为小数，1必须放在中间
			return ones - bits.OnesCount(mask&0xAAAAAAAA) // TODO 取有效位 ...1010 1010
		} else { //ones == half+1
			//  1位大数，1必须放在两边
			return ones - bits.OnesCount(mask&0x55555555) // ...0101 0101
		}
	} else {
		if ones != half || cnt != half {
			return -1
		}
		// TODO: 两边1少，移动哪边
		cnt0 := bits.OnesCount(mask & 0xAAAAAAAA)
		cnt1 := ones - cnt0
		if cnt0 < cnt1 {
			return cnt0
		} else {
			return cnt1
		}
	}
}

func movesToChessboard(board [][]int) int {
	bL := len(board)
	// 棋盘的第一行与第一列
	maskRow, maskCol := 0, 0
	for i := 0; i < bL; i++ {
		maskRow |= board[0][i] << i // TODO
		maskCol |= board[i][0] << i
	}

	maskRowMirror := 1<<bL - 1 ^ maskRow // TODO 镜像
	maskColMirror := 1<<bL - 1 ^ maskCol
	rowCnt, colCnt := 0, 0
	for i := 0; i < bL; i++ {
		maskRowCur, maskColCur := 0, 0
		for j := 0; j < bL; j++ {
			maskRowCur |= board[i][j] << j
			maskColCur |= board[j][i] << j
		}

		// TODO
		if maskRowCur != maskRow && maskRowCur != maskRowMirror || // 检测每一行的状态是否合法
			maskColCur != maskCol && maskColCur != maskColMirror { // 检测每一列的状态是否合法
			return -1
		}
		if maskRowCur == maskRow {
			rowCnt++ // 记录与第一行相同的行数
		}
		if maskColCur == maskCol {
			colCnt++ // 记录与第一列相同的列数
		}
	}
	rowMoves := getMoves(uint(maskRow), rowCnt, bL)
	if rowMoves == -1 {
		return -1
	}
	colMoves := getMoves(uint(maskCol), colCnt, bL)
	if colMoves == -1 {
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
