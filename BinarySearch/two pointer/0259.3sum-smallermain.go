package two_pointer

import "sort"

// https://leetcode-cn.com/problems/3sum-smaller

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	numsL := len(nums)
	var sum int
	for i := 0; i < numsL; i++ {
		sum += twoSumSmaller(nums, i+1, target-nums[i])
	}
	return sum
}

func twoSumSmaller(nums []int, idxLeft, target int) int {
	var sum int
	idxRight := len(nums) - 1
	for idxLeft < idxRight {
		numCur := nums[idxLeft] + nums[idxRight]
		if numCur < target {
			sum += idxRight - idxLeft
			idxLeft++
		} else {
			idxRight--
		}
	}
	return sum
}
