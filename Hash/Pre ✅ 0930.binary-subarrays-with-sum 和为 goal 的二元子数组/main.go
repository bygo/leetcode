package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

// ❓ 和为 goal 的二元子数组
// ⚠️ 1 0 0 1 0 0 0 0 1
//  0 1 1 1 2 2 2 2 2 3

func numSubarraysWithSum(nums []int, goal int) int {
	var cnt, sum int
	sumMpCnt := map[int]int{}
	for _, num := range nums {
		sumMpCnt[sum]++ // 0
		sum += num
		cnt += sumMpCnt[sum-goal] // 前缀和 sum[j] - goal = sum[i]  包含自身的组合
	}
	return cnt
}
