package main

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays

// ❓ 无限翻转子数组 使两个数组相等

func canBeEqual(target []int, arr []int) bool {
	targetL := len(target)
	arrL := len(arr)
	if targetL != arrL {
		return false
	}
	numMpCnt := map[int]int{}
	for i := 0; i < targetL; i++ {
		numMpCnt[target[i]]++
		numMpCnt[arr[i]]--
	}

	for i := range numMpCnt {
		if numMpCnt[i] != 0 {
			return false
		}
	}
	return true
}
