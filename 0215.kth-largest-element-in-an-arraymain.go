package main

// https://leetcode.cn/problems/kth-largest-element-in-an-array

func findKthLargest(nums []int, k int) int {
	var cnts [20001]int
	for _, num := range nums {
		cnts[num+10000]++
	}

	n := len(nums) - k
	for num, cnt := range cnts {
		n -= cnt
		if n < 0 {
			return num - 10000
		}
	}
	return -1
}
