package main

// https://leetcode.cn/problems/largest-unique-number

// ❓ 最大的唯一值

func largestUniqueNumber(nums []int) int {
	// 计数
	numMpCnt := [1001]int{}
	for _, num := range nums {
		numMpCnt[num]++
	}

	// 倒序
	for num := 1000; 0 <= num; num-- {
		if numMpCnt[num] == 1 {
			return num
		}
	}
	return -1
}
