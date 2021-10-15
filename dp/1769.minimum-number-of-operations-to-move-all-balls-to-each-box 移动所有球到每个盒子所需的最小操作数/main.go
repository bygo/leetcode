package main

// https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

// Pre
func minOperations(boxes string) []int {
	var cnt, cur int
	var l = len(boxes)
	var res = make([]int, l)
	for i := 0; i < l; i++ {
		cnt += cur
		res[i] += cnt
		if boxes[i] == '1' {
			cur++
		}
	}

	cnt, cur = 0, 0
	for i := l - 1; 0 < i; i-- {
		cnt += cur
		res[i] += cnt
		if boxes[i] == '1' {
			cur++
		}
	}

	return res
}
