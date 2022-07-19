package main

import "sort"

// https://leetcode.cn/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums) //排序，保证从小到大
	var one, two, left, right int
	l := len(nums)
	for ; one < l; one++ {
		if one > 0 && nums[one] == nums[one-1] { //如果值相同，直接跳过，因为后续重复的值，所有的匹配情况都是前面值的子集
			continue
		}
		two = one + 1
		for ; two < l-1; two++ {
			if two > one+1 && nums[two] == nums[two-1] {
				continue
			}
			left = two + 1
			right = l - 1
			for left < right {
				sum := nums[one] + nums[two] + nums[left] + nums[right]
				if target < sum { //太大，指针right往左移动
					right--
				} else if sum < target { //太小，指针left往右移动
					left++
				} else { //成功匹配后，移除 `right前序`或者`left后序`所有重复值
					res = append(res, []int{
						nums[one], nums[two], nums[left], nums[right],
					})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					right--
					left++
				}
			}
		}
	}
	return res
}
