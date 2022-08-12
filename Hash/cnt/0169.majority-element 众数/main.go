package main

// https://leetcode.cn/problems/majority-element

// ❓ 超过一半的数

func majorityElement(nums []int) int {
	var numMode, cnt int
	for _, num := range nums {
		if numMode == num {
			cnt++
		} else {
			if cnt == 0 {
				numMode = num
				cnt = 1
			} else {
				cnt--
			}
		}
	}
	return numMode
}
