package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

// ❓ 出现次数cnt等于值val的最大值

func findLucky(arr []int) int {
	// 计数
	numMpCnt := [501]int{}
	for i := range arr {
		numMpCnt[arr[i]]++
	}

	// 倒序
	for num := 500; 0 < num; num-- {
		if numMpCnt[num] == num {
			return num
		}
	}
	return -1
}
