package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

// ❓sum(相同 num) == num 的 最大数

func findLucky(arr []int) int {
	numMpCnt := [501]int{}
	for i := range arr {
		numMpCnt[arr[i]]++
	}

	for num := 500; 0 < num; num-- {
		if numMpCnt[num] == num {
			return num
		}
	}
	return -1
}
