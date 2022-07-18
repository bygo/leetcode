package main

// https://leetcode.cn/problems/line-reflection

// ❓ 直线镜像

func isReflected(points [][]int) bool {
	xMpYMpBool := map[int]map[int]bool{}
	xMax, xMin := -1<<63, 1<<63-1

	// 找出最大范围 并统计
	for _, point := range points {
		x := point[0]
		y := point[1]
		if xMax < x {
			xMax = x
		}
		if x < xMin {
			xMin = x
		}
		if xMpYMpBool[x] == nil {
			xMpYMpBool[x] = map[int]bool{}
		}
		xMpYMpBool[x][y] = true
	}

	// 使所有平行，必须找出中垂线

	for _, point := range points {
		x := xMax + xMin - point[0]
		y := point[1]
		// x 镜像不存在 或者 不平行
		if xMpYMpBool[x] == nil || !xMpYMpBool[x][y] {
			return false
		}
	}
	return true
}
