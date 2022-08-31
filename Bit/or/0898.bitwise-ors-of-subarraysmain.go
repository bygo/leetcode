package main

// https://leetcode.cn/problems/bitwise-ors-of-subarrays

func subarrayBitwiseORs(nums []int) int {
	maskMp := map[int]struct{}{}
	for idx, num := range nums {
		maskMp[num] = struct{}{}
		mask := 0
		for j := idx - 1; 0 <= j; j-- {
			if mask|num == mask { // TODO
				break
			}
			mask |= nums[j]
			maskMp[mask|num] = struct{}{}
		}
	}
	return len(maskMp)
}
