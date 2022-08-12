package main

func mul(x, y int) int {
	var numRes int
	var cnt int
	// 符号转换
	if x < 0 {
		cnt++
		x = -x
	}
	if y < 0 {
		cnt++
		y = -y
	}

	for 0 < y {
		// 按位相加
		if 1 == y&1 {
			numRes += x
		}
		// 减少位
		y >>= 1
		// 翻倍
		x <<= 1
	}
	if cnt%2 == 1 {
		return -numRes
	}
	return numRes
}
