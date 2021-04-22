package main

//Title: Minimum Number of Operations to Move All Balls to Each Box
//Link: https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

func minOperations(boxes string) []int {
	var res []int
	var counter, cur int
	var l = len(boxes) - 1
	for i := 0; i <= l; i++ {
		counter += cur
		res = append(res, counter)
		if boxes[i] == '1' {
			cur++
		}
	}
	counter, cur = 0, 0
	for i := l; 0 <= i; i-- {
		counter += cur
		res[i] += counter
		if boxes[i] == '1' {
			cur++
		}
	}

	return res
}
