package main

// https://leetcode-cn.com/problems/find-the-duplicate-number

// ❓ 寻找重复数
// ⚠️ n = 1 ~ numsL-1

func findDuplicate(nums []int) int {
	numsL := len(nums)
	posMax := 17
	// 确认最高位
	for (numsL-1)>>posMax == 0 {
		posMax--
	}

	numRes := 0
	// 按位确认
	for pos := 0; pos < posMax; pos++ {
		var cntNum, cntIdx int
		for idx, num := range nums {
			// 1 ~ numsL - 1
			// 重复的话，对应pos 就会多出更多的1
			if 1 == num>>pos&1 {
				cntNum++
			}
			// 0 1 ~ numsL - 1 ，用idx，只为了求范围内所有可能的1
			if 1 == idx>>pos&1 {
				cntIdx++
			}
		}

		// 重复位
		if cntIdx < cntNum {
			numRes |= 1 << pos
		}
	}
	return numRes
}
