package main

// https://leetcode-cn.com/problems/line-reflection

func isReflected(points [][]int) bool {
	m := map[int]map[int]bool{}
	max, min := -1<<63, 1<<63-1

	for _, point := range points {
		if max < point[0] {
			max = point[0]
		}
		if point[0] < min {
			min = point[0]
		}
		if m[point[0]] == nil {
			m[point[0]] = map[int]bool{}
		}
		m[point[0]][point[1]] = true
	}
	for _, point := range points {
		x := max + min - point[0]
		// x 镜像不存在 或者 不平行
		if m[x] == nil || !m[x][point[1]] {
			return false
		}
	}
	return true
}
