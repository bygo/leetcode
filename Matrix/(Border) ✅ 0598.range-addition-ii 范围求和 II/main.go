package main

// https://leetcode-cn.com/problems/range-addition-ii

func maxCount(m int, n int, ops [][]int) int {
	var x = n
	var y = m

	var cnt = len(ops)
	for i := 0; i < cnt; i++ {
		if ops[i][1] < x {
			x = ops[i][1]
		}
		if ops[i][0] < y {
			y = ops[i][0]
		}
	}

	return x * y
}
