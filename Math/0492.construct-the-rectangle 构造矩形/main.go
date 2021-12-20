package main

// https://leetcode-cn.com/problems/construct-the-rectangle

func constructRectangle(area int) []int {
	if area == 1 {
		return []int{1, 1}
	}
	l, r := 1, area
	for l < r {
		mid := (l + r) >> 1
		cur := mid * mid
		if cur <= area { // sqrt
			l = mid + 1
		} else {
			r = mid
		}
	}

	l--
	for 0 < area%l {
		l--
	}
	return []int{area / l, l}
}
