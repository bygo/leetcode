package main

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist

// ❓是否存在两倍数

func checkIfExist(arr []int) bool {
	numMpCnt := map[int]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	for num := range numMpCnt {
		if 0 < numMpCnt[num*2] {
			if num != 0 || 2 <= numMpCnt[num] {
				return true
			}
		}
	}
	return false
}
