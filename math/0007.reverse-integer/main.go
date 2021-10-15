package main

func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if y < -1<<31 || 1<<31-1 < y {
			return 0
		}
		x /= 10
	}
	return y
}
