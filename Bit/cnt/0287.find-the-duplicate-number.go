package main

// https://leetcode.cn/problems/find-the-duplicate-number

// 1 <= n <= 10^5
// nums.length == n + 1
// 1 <= nums[i] <= n

func findDuplicate(nums []int) int {
	numsL := len(nums)
	idxMax := 17
	numMax := numsL - 1
	for numMax>>idxMax == 0 { // TODO
		idxMax--
	}

	numRes := 0
	for idx := 0; idx <= idxMax; idx++ {
		var cntNum, cntIdx int // TODO
		for idxNum, num := range nums {
			// 1 ~ numsL-1
			if 1 == num>>idx&1 {
				cntNum++
			}
			// 0 1 ~ numsL-1
			if 1 == idxNum>>idx&1 {
				cntIdx++
			}
		}

		if cntIdx < cntNum {
			numRes |= 1 << idx
		}
	}
	return numRes
}
