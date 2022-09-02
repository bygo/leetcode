package main

// https://leetcode.cn/problems/bitwise-ors-of-subarrays

func subarrayBitwiseORs(nums []int) int {
	maskMp := map[int]struct{}{}
	for idx, num := range nums {
		maskMp[num] = struct{}{}
		mask := 0
		for j := idx - 1; 0 <= j; j-- {
			if mask|num == mask { // TODO 前置包含num,num 对mask 不起作用
				break
			}
			mask |= nums[j]
			maskMp[mask|num] = struct{}{} // TODO num能让mask发生变化
		}
	}
	return len(maskMp)
}
