package two_pointer

import "sort"

// https://leetcode.cn/problems/3sum-smaller

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	numsL := len(nums)
	var cnt int
	for i := 0; i < numsL; i++ {
		cnt += twoSumSmaller(nums, i+1, target-nums[i])
	}
	return cnt
}

func twoSumSmaller(nums []int, lo, target int) int {
	var cnt int
	hi := len(nums) - 1
	for lo < hi {
		numCur := nums[lo] + nums[hi]
		if target <= numCur {
			hi--
		} else if numCur < target {
			cnt += hi - lo
			lo++
		}
	}
	return cnt
}
