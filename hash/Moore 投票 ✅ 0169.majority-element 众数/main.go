package main

// https://leetcode-cn.com/problems/majority-element

func majorityElement(nums []int) int {
	var num, cnt int
	for i := range nums {
		if num == nums[i] {
			cnt++
		} else {
			if cnt == 0 {
				num = nums[i]
				cnt = 1
			} else {
				cnt--
			}
		}
	}
	return num
}
