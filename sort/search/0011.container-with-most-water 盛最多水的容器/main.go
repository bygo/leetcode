package main

// https://leetcode-cn.com/problems/container-with-most-water/

func maxArea(height []int) int {
	var cur, max, l int
	r := len(height) - 1
	for l < r {
		if height[l] < height[r] { //最小高度乘以宽度，并且移动一位
			cur = height[l] * (r - l)
			l++
		} else {
			cur = height[r] * (r - l)
			r--
		}
		if max < cur {
			max = cur
		}

	}
	return max
}
