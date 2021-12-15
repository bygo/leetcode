package main

// https://leetcode-cn.com/problems/sum-of-unique-elements

// ❓ 所有唯一元素的和

func sumOfUnique(nums []int) int {
	var numMpCnt = map[int]int{}
	var res int
	for _, num := range nums {
		if numMpCnt[num] == 0 {
			res += num
			numMpCnt[num] = 1
		} else if numMpCnt[num] == 1 {
			res -= num
			numMpCnt[num] = 2
		}
	}
	return res
}
