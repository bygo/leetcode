package main

// https://leetcode-cn.com/problems/destination-city

func destCity(paths [][]string) string {
	m := map[string]struct{}{}
	for _, path := range paths {
		m[path[0]] = struct{}{}
	}
	for _, path := range paths {
		_, ok := m[path[1]]
		if !ok {
			return path[1]
		}
	}
	return ""
}

func gridGame(grid [][]int) int64 {
	l1 := len(grid)
	l2 := len(grid[0])

	var path = make([][]int, l1)
	for i := range path {
		path[i] = make([]int, l2)
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		nx1, ny1 := x+1, y
		nx2, ny2 := x, y+1
		if nx1 == l1 && ny1 == l2 {

		}
		if nx1 <= l1 && ny1 <= l2 {

		}

		if nx2 <= l1 && ny2 <= l2 {

		}
	}
}

