package main

// https://leetcode.cn/problems/counting-elements

// ❓ 存在后置元素的元素个数

func countElements(arr []int) int {
	numMpCnt := map[int]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	var cnt int
	for num := range numMpCnt {
		if 0 < numMpCnt[num+1] {
			cnt += numMpCnt[num]
		}
	}
	return cnt
}
