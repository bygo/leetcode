package main

// https://leetcode.cn/problems/relative-sort-array

// ❓ 数组A 遵循 数组B 正序

func relativeSortArray(arr1 []int, arr2 []int) []int {
	numMpCnt := [1001]int{}
	for _, num := range arr1 {
		numMpCnt[num]++
	}

	var numsSort []int

	// 按数组B 排序
	for _, num := range arr2 {
		for 0 < numMpCnt[num] {
			numsSort = append(numsSort, num)
			numMpCnt[num]--
		}
	}

	// 未出现过的 正序
	for num := range numMpCnt {
		for 0 < numMpCnt[num] {
			numsSort = append(numsSort, num)
			numMpCnt[num]--
		}
	}
	return numsSort
}
