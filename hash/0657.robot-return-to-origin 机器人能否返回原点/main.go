package main

// Link: https://leetcode-cn.com/problems/robot-return-to-origin

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := range moves {
		switch moves[i] {
		case 'U':
			y--
		case 'D':
			y++
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
