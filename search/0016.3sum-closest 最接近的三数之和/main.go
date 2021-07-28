package main

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/3sum-closest/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) //排序，保证从小到大
	var one, left, right int
	var l = len(nums)
	res := nums[0] + nums[1] + nums[2]
	for ; one < l; one++ {
		if 0 < one && nums[one] == nums[one-1] { //如果值相同，直接跳过，因为后续重复的值，所有的匹配情况都是前面值的子集
			continue
		}
		left = one + 1
		right = len(nums) - 1
		for left < right {
			sum := nums[one] + nums[left] + nums[right]
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
