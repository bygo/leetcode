package main

// https://leetcode-cn.com/problems/majority-element

func majorityElement(nums []int) int {
	var cur, cnt int
	for _, num := range nums {
		if cnt == 0 {
			cur = num
		}
		if cur == num {
			cnt++
		} else {
			cnt--
		}
	}
	return cur
}
