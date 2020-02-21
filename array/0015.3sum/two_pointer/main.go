package main

import (
	"sort"
)

/**
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

 */

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums) //排序，保证从小到大
	var one, left, right int
	var l = len(nums)
	for ; one < l; one++ {
		if nums[one] > 0 { //最小值大于0，三数必定大于0,
			break
		}
		if one > 0 && nums[one] == nums[one-1] { //如果值相同，直接跳过，因为后续重复的值，所有的匹配情况都是前面值的子集
			continue
		}
		left = one + 1
		right = len(nums) - 1
		for left < right {
			sum := nums[one] + nums[left] + nums[right]
			if sum > 0 { //太大，指针right往左移动
				right--
			} else if sum < 0 { //太小，指针left往右移动
				left++
			} else { //成功匹配后，移除 `right前序`或者`left后序`所有重复值
				res = append(res, []int{
					nums[one], nums[left], nums[right],
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
	return res
}

/**
思路:
1.遍历+双指针
 */
