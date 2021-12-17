package main

import "sort"

// https://leetcode-cn.com/problems/WqXACV/

// ❓ 完成一半题目最少题型
// ⚠️ 2*N 道题目
func halfQuestions(questions []int) int {
	// 类型计数
	typMpCnt := make([]int, 1001)
	half := len(questions) / 2
	for i := range questions {
		typMpCnt[questions[i]]++
	}

	// 类型排序
	sort.Slice(typMpCnt, func(i, j int) bool {
		return typMpCnt[j] < typMpCnt[i]
	})

	// 最少类型
	var cntTyp int
	for i := range typMpCnt {
		if half <= 0 {
			return cntTyp
		}
		half -= typMpCnt[i]
		cntTyp++
	}
	return 10086
}
