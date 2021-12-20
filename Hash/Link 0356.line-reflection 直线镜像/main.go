package main

// https://leetcode-cn.com/problems/line-reflection

func isReflected(points [][]int) bool {
	xMpYMpBool := map[int]map[int]bool{}
	max, min := -1<<63, 1<<63-1

	// 找出最大范围 并统计
	for _, point := range points {
		x := point[0]
		y := point[1]
		if max < x {
			max = x
		}
		if x < min {
			min = x
		}
		if xMpYMpBool[x] == nil {
			xMpYMpBool[x] = map[int]bool{}
		}
		xMpYMpBool[x][y] = true
	}

	// 使所有平行，必须找出中垂线

	for _, point := range points {
		x := max + min - point[0]
		y := point[1]
		// x 镜像不存在 或者 不平行
		if xMpYMpBool[x] == nil || !xMpYMpBool[x][y] {
			return false
		}
	}
	return true
}
