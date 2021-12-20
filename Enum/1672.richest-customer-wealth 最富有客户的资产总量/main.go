package main

// https://leetcode-cn.com/problems/richest-customer-wealth

func maximumWealth(accounts [][]int) int {
	var max int
	for _, account := range accounts {
		var cur int
		for _, v := range account {
			cur += v
		}
		if max < cur {
			max = cur
		}
	}
	return max
}
