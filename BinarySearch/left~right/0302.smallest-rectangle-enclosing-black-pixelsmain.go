package main

// https://leetcode.cn/problems/smallest-rectangle-enclosing-black-pixels

func minArea(image [][]byte, y int, x int) int {
	rowL, colL := len(image), len(image[0])
	left := searchCols(image, 0, x, 0, rowL, true)
	right := searchCols(image, x+1, colL, 0, rowL, false)
	top := searchRows(image, left, right, 0, y, true)
	bottom := searchRows(image, left, right, y+1, rowL, false)
	return (right - left) * (bottom - top)
}

func searchRows(image [][]byte, left, right, top, bottom int, whiteToBlack bool) int {
	lo, hi := top, bottom
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		colK := left
		for colK < right && image[mid][colK] == '0' {
			colK++
		}

		hasBlack := colK < right
		if hasBlack == whiteToBlack {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func searchCols(image [][]byte, left, right, top, bottom int, whiteToBlack bool) int {
	lo, hi := left, right
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		rowK := top
		for rowK < bottom && image[rowK][mid] == '0' {
			rowK++
		}
		hasBlack := rowK < bottom
		if hasBlack == whiteToBlack {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
