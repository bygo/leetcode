package main

// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game

const (
	A       = "A"
	B       = "B"
	Pending = "Pending"
	Draw    = "Draw"
)

func tictactoe(moves [][]int) string {
	l1 := len(moves)

	m := [8]int{}

	for i := l1 - 1; 0 <= i; i -= 2 {
		m[moves[i][0]]++
		m[moves[i][1]+3]++
		if moves[i][0] == moves[i][1] {
			m[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			m[7]++
		}
	}

	for i := range m {
		if m[i] != 3 {
			continue
		}
		if l1%2 == 1 {
			return A
		} else {
			return B
		}
	}
	if l1 < 9 {
		return Pending
	}

	return Draw
}
