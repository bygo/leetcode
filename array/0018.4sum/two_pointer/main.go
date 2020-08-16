package main

import "sort"

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

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

/**
思路：
N数之和=N个指针

todo:如果还有N数之和，再封装函数。
*/
