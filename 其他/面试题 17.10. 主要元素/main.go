package main

// https://leetcode-cn.com/problems/find-majority-element-lcci/

func majorityElement(nums []int) int {
	l1 := len(nums)
	candidate := -1
	cnt := 0
	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}
		if num == candidate {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, num := range nums {
		if num == candidate {
			cnt++
		}
	}
	if l1 < cnt*2 {
		return candidate
	}
	return -1
}

func majorityElement(nums []int) int {
	l1 := len(nums)
	candidate := -1
	cnt := 0
	for _, num := range nums {
		if cnt == 0 {
			candidate = num
		}
		if candidate == num {
			cnt++
		} else {
			cnt--
		}
	}

	cnt = 0
	for _, num := range nums {
		if num == candidate {
			cnt++
		}
	}
	if l1/2 < cnt {
		return candidate
	}
	return -1
}
