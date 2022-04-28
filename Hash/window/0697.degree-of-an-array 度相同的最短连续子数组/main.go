package main

// https://leetcode-cn.com/problems/degree-of-an-array

// ❓ 度相同的最短连续子数组
// ⚠️ 度 指 数组 其中一元素频数的最大值。

type node struct {
	cnt   int
	left  int
	right int
}

// 最短 且 最大的度必定横跨首尾

func findShortestSubArray(nums []int) int {
	numMpNode := map[int]*node{}
	for idx, num := range nums {
		_, ok := numMpNode[num]
		if ok {
			// 已存在，计算右边界
			numMpNode[num].right = idx
			numMpNode[num].cnt++
		} else {
			numMpNode[num] = &node{
				cnt:   1,
				left:  idx,
				right: idx,
			}
		}
	}

	var cntMax int
	var shortestL int
	for _, node := range numMpNode {
		if cntMax < node.cnt {
			// 频率大于之前 重算
			cntMax = node.cnt
			shortestL = node.right - node.left + 1
		} else if cntMax == node.cnt {

			// 频率相同时，计算最短长度
			curL := node.right - node.left + 1
			if curL < shortestL {
				shortestL = curL
			}
		}
	}
	return shortestL
}
