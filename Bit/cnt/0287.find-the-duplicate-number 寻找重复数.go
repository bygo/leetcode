package main

// https://leetcode-cn.com/problems/find-the-duplicate-number

// ❓ 寻找重复数
// ⚠️ n = 1 ~ numsL-1

func findDuplicate(nums []int) int {
	numsL := len(nums)
	idxMax := 17
	numMax := numsL - 1
	// 确认最高位
	for numMax>>idxMax == 0 {
		idxMax--
	}

	numRes := 0
	// 按位确认
	for idx := 0; idx <= idxMax; idx++ {
		var cntNum, cntIdx int
		for i, num := range nums {
			// 1 ~ numsL - 1
			// 重复的话，对应pos 就会多出更多的1
			if 1 == num>>idx&1 {
				cntNum++
			}
			// 0忽略 1 ~ numsL - 1 ，用i，只为了求范围内所有可能的1
			if 1 == i>>idx&1 {
				cntIdx++
			}
		}

		// 当前是否有 重复位
		if cntIdx < cntNum {
			// 加入
			numRes |= 1 << idx
		}
	}
	return numRes
}
