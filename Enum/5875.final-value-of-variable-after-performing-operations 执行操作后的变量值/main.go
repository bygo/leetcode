package main

// https://leetcode.cn/problems/final-value-of-variable-after-performing-operations

func finalValueAfterOperations(operations []string) int {
	var res int
	for i := range operations {
		if operations[i][1] == '+' {
			res++
		} else if operations[i][1] == '-' {
			res--
		}
	}
	return res
}
