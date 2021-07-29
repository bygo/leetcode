package main

// https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

func minOperations(boxes string) []int {
	var res []int
	var cnt, cur int
	var l = len(boxes) - 1
	for i := 0; i <= l; i++ {
		cnt += cur
		res = append(res, cnt)
		if boxes[i] == '1' {
			cur++
		}
	}

	cnt, cur = 0, 0
	for i := l; 0 <= i; i-- {
		cnt += cur
		res[i] += cnt
		if boxes[i] == '1' {
			cur++
		}
	}

	return res
}
