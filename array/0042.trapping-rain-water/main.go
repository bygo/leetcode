package main

// Link:https://leetcode-cn.com/problems/trapping-rain-water/

func trap(h []int) int {
	var l, r, lMax, rMax, res int
	r = len(h) - 1
	for l < r {
		if h[l] < h[r] { // 因为 h[l] < h[r] 右挡板
			if lMax < h[l] {
				lMax = h[l] // 左挡板 最高高度
			} else {
				res += lMax - h[l]
			}
			l++
		} else { // 因为 h[r] <= h[l] 左挡板
			if rMax < h[r] {
				rMax = h[r] // 右挡板 最高高度
			} else {
				res += rMax - h[r]
			}
			r--
		}
	}
	return res
}
