package main

func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left] //设置左边最高柱子
			} else {
				res += leftMax - height[left] //右边必定有柱子挡水，所以，遇到所有值小于leftMax的，全部加入水池
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right] //设置右边最高柱子
			} else {
				res += rightMax - height[right] //左边必定有柱子挡水，所以，遇到所有值小于rightMax的，全部加入水池
			}
			right--
		}
	}
	return res
}
