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
	l1 := len(moves)

	perMpCnt := [8]int{}

	for i := l1 - 1; 0 <= i; i -= 2 {
		row := moves[i][0]
		col := moves[i][1] + 3
		perMpCnt[row]++
		perMpCnt[col]++
		if moves[i][0] == moves[i][1] {
			perMpCnt[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			perMpCnt[7]++
		}
	}

	for i := range perMpCnt {
		if perMpCnt[i] != 3 {
			continue
		}
		if l1%2 == 1 {
			return A
		}
		return B
	}
	if l1 < 9 {
		return Pending
	}

	return Draw
}
