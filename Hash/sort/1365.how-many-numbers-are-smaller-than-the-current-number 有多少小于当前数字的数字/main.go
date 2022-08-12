package main

// https://leetcode.cn/problems/how-many-numbers-are-smaller-than-the-current-number/

// ❓ 有多少小于当前数字的数字
// ⚠️ 0 <= nums[i] <= 100

func smallerNumbersThanCurrent(nums []int) []int {
	var numMpCnt [101]int
	numsGt := make([]int, len(nums))

	for _, num := range nums {
		numMpCnt[num]++
	}

	// 前缀和
	for num := 1; num < 101; num++ {
		numMpCnt[num] += numMpCnt[num-1]
	}

	// 转换
	for i, num := range nums {
		if num != 0 {
			numsGt[i] = numMpCnt[num-1]
		}
	}
	return numsGt
}
