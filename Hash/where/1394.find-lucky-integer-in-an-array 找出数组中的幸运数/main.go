package main

// https://leetcode.cn/problems/find-lucky-integer-in-an-array

// ❓ 出现次数cnt等于值num的最大值
// ⚠️ 1 <= arr[i] <= 500
func findLucky(arr []int) int {
	// 计数
	numMpCnt := [501]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	// 倒序
	for num := 500; 0 < num; num-- {
		if numMpCnt[num] == num {
			return num
		}
	}
	return -1
}
