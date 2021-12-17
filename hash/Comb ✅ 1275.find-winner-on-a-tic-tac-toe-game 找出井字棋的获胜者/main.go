package main

// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game

// ❓ 3x3 井字棋

const (
	A       = "A"
	B       = "B"
	Pending = "Pending"
	Draw    = "Draw"
)

func tictactoe(moves [][]int) string {
	movesL := len(moves)

	combMpCnt := [8]int{}

	for i := movesL - 1; 0 <= i; i -= 2 {
		row := moves[i][0]
		col := moves[i][1] + 3
		combMpCnt[row]++
		combMpCnt[col]++
		if moves[i][0] == moves[i][1] {
			combMpCnt[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			combMpCnt[7]++
		}
	}

	for i := range combMpCnt {
		if combMpCnt[i] != 3 {
			continue
		}
		if movesL%2 == 1 {
			return A
		}
		return B
	}
	if movesL < 9 {
		return Pending
	}

	return Draw
}
