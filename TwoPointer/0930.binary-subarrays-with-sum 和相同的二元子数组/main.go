package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

// ❓ 和为 goal 的二元子数组
// ⚠️ 1 0 0 1 0 0 0 0 1
//  0 1 1 1 2 2 2 2 2 3

func numSubarraysWithSum(nums []int, goal int) int {
	var cnt int
	idx1, idx2 := 0, 0
	sum1, sum2 := 0, 0
	for right, num := range nums {
		sum1 += num
		for idx1 <= right && goal < sum1 {
			sum1 -= nums[idx1]
			idx1++
		}
		sum2 += num
		for idx2 <= right && goal <= sum2 {
			sum2 -= nums[idx2]
			idx2++
		}
		// 索引差
		cnt += idx2 - idx1
	}
	return cnt
}

// ☝️ Hash
