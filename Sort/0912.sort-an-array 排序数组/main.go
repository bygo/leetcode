package main

// https://leetcode.cn/problems/sort-an-array

// ❓ 计数排序

func sortArray(nums []int) []int {
	numMpCnt := [100001]int{}

	for _, num := range nums {
		numMpCnt[num+50000]++
	}

	var numsSort []int
	for num, cnt := range numMpCnt {
		for i := 0; i < cnt; i++ {
			numsSort = append(numsSort, num-50000)
		}
	}
	return numsSort
}
