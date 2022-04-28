package main

// https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box

func minOperations(boxes string) []int {
	bL := len(boxes)
	var numsMax = make([]int, bL)
	var cntPre, cntSuf int
	var cntLeft, cntRight int
	for i := 0; i < bL; i++ {
		// 当前元素答案
		numsMax[i] += cntLeft
		// 前缀
		cntPre += int(boxes[i] - '0')
		// 每移动一次，都需加一次
		cntLeft += cntPre

		numsMax[bL-i-1] += cntRight
		cntSuf += int(boxes[bL-i-1] - '0')
		cntRight += cntSuf
	}
	return numsMax
}
