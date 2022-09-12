package main

// https://leetcode.cn/problems/binary-subarrays-with-sum

// ❓ 和为 goal 的二元子数组
// ⚠️ 1 0 0 1 0 0 0 0 1 0 0 0
//  0 1 1 1 2 2 2 2 2 3 3 3 3

// ⚠️ idxLeft2 和 idxLeft1 的差 指代这个整体有多少种组合,也就是前置0有多少个

func numSubarraysWithSum(nums []int, goal int) int {
	var cnt, idxLeft1, idxLeft2, sum1, sum2 int
	for idxRight, num := range nums {
		// 寻找一个合法的整体
		sum1 += num
		for idxLeft1 <= idxRight && goal < sum1 {
			sum1 -= nums[idxLeft1]
			idxLeft1++
		}

		// 0与合法的整体构成组合, idxLeft2为整体中第一个1的下一个指针
		sum2 += num
		for idxLeft2 <= idxRight && goal <= sum2 {
			sum2 -= nums[idxLeft2]
			idxLeft2++
		}

		// 遇到1 变动指针
		// 遇到0 组合
		// 整体本身也算一种组合
		cnt += idxLeft2 - idxLeft1
	}
	return cnt
}

// 双指针 一个合法的整体 寻找组合
// 前缀树 寻找 合法的组合
