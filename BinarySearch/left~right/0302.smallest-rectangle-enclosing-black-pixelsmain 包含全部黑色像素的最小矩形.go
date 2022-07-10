package main

// https://leetcode-cn.com/problems/smallest-rectangle-enclosing-black-pixels

func minArea(image [][]byte, x int, y int) int {
	rows, cols := len(image), len(image[0])
	left := searchCols(image, 0, y, 0, rows, true)
	right := searchCols(image, y+1, cols, 0, rows, false)
	top := searchRows(image, 0, x, left, right, true)
	bottom := searchRows(image, x+1, rows, left, right, false)
	return (right - left) * (bottom - top)
}

func searchCols(image [][]byte, lo, hi, top, bottom int, findLeft bool) int {
	for lo < hi {
		rowK := top
		mid := int(uint(lo+hi) >> 1)
		for rowK < bottom && image[rowK][mid] == '0' {
			rowK++
		}
		// rowK < bottom 等价于 找到黑色

		// 如果找left ，就必须 缩hi
		// 如果找right，就必须 缩lo

		// 其他情况 取反
		if rowK < bottom == findLeft {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func searchRows(image [][]byte, lo, hi, left, right int, findTop bool) int {
	for lo < hi {
		colK := left
		mid := int(uint(lo+hi) >> 1)
		for colK < right && image[mid][colK] == '0' {
			colK++
		}
		// colK < right 等价于 找到黑色
		// 如果找top ，  就必须 缩hi
		// 如果找bottom，就必须 缩lo

		// 其他情况 取反
		if colK < right == findTop {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
