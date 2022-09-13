package main

import "sort"

// https://leetcode.cn/problems/mean-of-array-after-removing-some-elements

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	arrL := len(arr)
	sum := 0
	for _, x := range arr[arrL/20 : 19*arrL/20] {
		sum += x
	}
	return float64(sum*10) / float64(arrL*9)
}
