package main

// https://leetcode.cn/problems/binary-subarrays-with-sum

// ❓ 和为 goal 的二元子数组
// ⚠️ 1 0 0 1 0 0 0 0 1 0 0 0 1
//  0 1 1 1 2 2 2 2 2 3 3 3 3 4

func numSubarraysWithSum(nums []int, goal int) int {
	var cnt, sum int
	sumMpCnt := map[int]int{}
	for _, num := range nums {
		sumMpCnt[sum]++ // 0
		sum += num
		// 实际在寻找前置零的个数，cnt = N（前置零） + 1（自身）
		cnt += sumMpCnt[sum-goal] // 前缀和 sum[j] - goal = sum[i]  只包含自身的组合
	}
	return cnt
}
