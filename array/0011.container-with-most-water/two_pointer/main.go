package main

/**
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
链接：https://leetcode-cn.com/problems/container-with-most-water


示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/

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
		if cur > max {
			max = cur
		}

	}
	return max
}

/**
思路：
1.双指针，移动最小值的指针，因为宽度减小，必须除去最短的值,才能获得最大值
*/
