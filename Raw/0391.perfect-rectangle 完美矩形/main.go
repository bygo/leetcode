package main

// https://leetcode-cn.com/problems/perfect-rectangle
type point struct{ x, y int }

func isRectangleCover(rectangles [][]int) bool {
	area, minX, minY, maxX, maxY := 0, rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	cnt := map[point]int{}
	for _, rect := range rectangles {
		x, y, a, b := rect[0], rect[1], rect[2], rect[3]
		area += (a - x) * (b - y)

		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, a)
		maxY = max(maxY, b)

		cnt[point{x, y}]++
		cnt[point{x, b}]++
		cnt[point{a, y}]++
		cnt[point{a, b}]++
	}

	if area != (maxX-minX)*(maxY-minY) ||
		cnt[point{minX, minY}] != 1 ||
		cnt[point{minX, maxY}] != 1 ||
		cnt[point{maxX, minY}] != 1 ||
		cnt[point{maxX, maxY}] != 1 {
		return false
	}

	delete(cnt, point{minX, minY})
	delete(cnt, point{minX, maxY})
	delete(cnt, point{maxX, minY})
	delete(cnt, point{maxX, maxY})

	for _, c := range cnt {
		if c != 2 && c != 4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
