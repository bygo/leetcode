package main

// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game

func tictactoe(moves [][]int) string {
	top := len(moves) - 1

	m := [8]int{}

	for i := top; 0 <= i; i -= 2 {
		m[moves[i][0]]++
		m[moves[i][1]+3]++
		if moves[i][0] == moves[i][1] {
			m[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			m[7]++
		}
	}

	l1 := len(moves)
	var res string
	if l1%2 == 0 {
		res = "B"
	} else {
		res = "A"
	}
	for i := range m {
		if m[i] == 3 {
			return res
		}
	}
	if l1 < 9 {
		return "Pending"
	}

	return "Draw"
}
