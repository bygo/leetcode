package main

import (
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 */

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) //排序，保证从小到大
	var left, right int
	res := nums[0] + nums[1] + nums[2]
	for k, one := range nums {
		if 0 < k && nums[k] == nums[k-1] { //如果值相同，直接跳过，因为后续重复的值，所有的匹配情况都是前面值的子集
			continue
		}
		left = k + 1
		right = len(nums) - 1
		for left < right {
			sum := one + nums[left] + nums[right]
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
			if target < sum { //太大，指针right往左移动
				right--
			} else if sum < target { //太小，指针left往右移动
				left++
			} else {
				return sum
			}

		}
	}
	return res
}

/**
思路:
1.与 threeSum 解法一样
 */
