package main

// https://leetcode.cn/problems/count-equal-and-divisible-pairs-in-an-array

// ❓ 统计数组中相等且可以被整除的数对

func countPairs(nums []int, k int) int {
	numMpIdxes := map[int][]int{}
	for idx, num := range nums {
		numMpIdxes[num] = append(numMpIdxes[num], idx)
	}
	var cntRes int
	for _, idxes := range numMpIdxes {
		for i, idx := range idxes {
			for _, idx2 := range idxes[i+1:] {
				if (idx*idx2)%k == 0 {
					cntRes++
				}
			}
		}
	}
	return cntRes
}
