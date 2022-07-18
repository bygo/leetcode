package main

// https://leetcode.cn/problems/trapping-rain-water/

// L and R go to the highest point together

func trap(h []int) int {
	var l, lMax, rMax, res int
	var r = len(h) - 1
	for l < r {
		if h[l] < h[r] {
			if lMax < h[l] {
				lMax = h[l]
			} else {
				res += lMax - h[l]
			}
			l++
		} else if h[r] <= h[l] {
			if rMax < h[r] {
				rMax = h[r]
			} else {
				res += rMax - h[r]
			}
			r--
		}
	}
	return res
}
