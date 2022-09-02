package main

// https://leetcode.cn/problems/find-the-duplicate-number

// 1 <= n <= 10^5
// nums.length == n + 1
// 1 <= nums[i] <= n

func findDuplicate(nums []int) int {
	numsL := len(nums)
	lo, hi := -1, 16
	numMax := 1<<numsL - 1
	for lo < hi { // TODO
		mid := int(uint(lo+hi+1) >> 1)
		if numMax>>mid&1 == 1 { // TODO
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	//hi := 17
	//for numMax>>hi == 0 { // TODO
	//	hi--
	//}

	numRes := 0
	for idx := 0; idx <= hi; idx++ {
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
