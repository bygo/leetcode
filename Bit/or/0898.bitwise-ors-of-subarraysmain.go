package main

// https://leetcode.cn/problems/bitwise-ors-of-subarrays

func subarrayBitwiseORs(nums []int) int {
	maskMp := map[int]struct{}{}
	numsL := len(nums)
	for idx, num := range nums {
		maskMp[num] = struct{}{}
		mask := 0
		for j := idx + 1; j < numsL; j++ {
			if mask|num == mask { // TODO 后置包含num的状态
				break
			}
			mask |= nums[j]
			maskMp[mask|num] = struct{}{}
		}
		//for j := idx - 1; 0 <= j; j-- {
		//	if mask|num == mask { // TODO 前置包含num的状态
		//		break
		//	}
		//	mask |= nums[j]
		//	maskMp[mask|num] = struct{}{} // TODO num能让mask发生变化
		//}
	}
	return len(maskMp)
}
