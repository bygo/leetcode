package main

func minTimeToVisitAllPoints(points [][]int) int {
	var res int
	l := len(points)
	for i := 1; i < l; i++ {
		x := abs(points[i][0] - points[i-1][0])
		y := abs(points[i][1] - points[i-1][1])
		res += max(x, y)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
