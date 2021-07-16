package main

//Title: Partitioning Into Minimum Number Of Deci-Binary Numbers
// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers

func minPartitions(n string) int {
	var max rune = -1
	for _, v := range n {
		v = v - 48
		if max < v {
			max = v
		}
	}
	return int(max)
}
