package main

// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers

func minPartitions(n string) int {
	var max rune = -1
	for _, v := range n {
		if max < v {
			max = v
		}
	}
	return int(max - '0')
}
