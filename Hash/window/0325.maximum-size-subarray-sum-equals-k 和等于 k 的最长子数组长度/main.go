package main

// https://leetcode.cn/problems/maximum-size-subarray-sum-equals-k

// ❓ 和等于K的最长子数组长度

func maxSubArrayLen(nums []int, k int) int {
	numMpIdx := map[int]int{0: -1}
	var sumPre, maxL, curL int
	for idxRight := range nums {
		// 寻找第一个出现时的idx

		sumPre += nums[idxRight]
		idxLeft, ok := numMpIdx[sumPre-k]
		if ok {
			// 不包括 idxLeft 所以不+1
			curL = idxRight - idxLeft
			if maxL < curL {
				maxL = curL
			}
		}

		// 不存在时加入
		_, ok = numMpIdx[sumPre]
		if !ok {
			numMpIdx[sumPre] = idxRight
		}
	}
	return maxL
}
