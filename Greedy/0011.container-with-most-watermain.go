package main

// https://leetcode.cn/problems/container-with-most-water

func maxArea(height []int) int {
	var cur, max int
	lo, hi := 0, len(height)-1
	for lo < hi {
		dist := hi - lo
		if height[lo] < height[hi] { // TODO 以小的为隔板
			cur = height[lo] * dist
			lo++
		} else {
			cur = height[hi] * dist
			hi--
		}

		// TODO 贪心
		if max < cur {
			max = cur
		}
	}
	return max
}
