package main

// https://leetcode.cn/problems/closest-binary-search-tree-value

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最接近的值

func closestValue(root *TreeNode, target float64) int {
	var numClosest float64 = 1<<63 - 1
	for root != nil {
		numCur := float64(root.Val)
		if abs(numCur-target) < abs(numClosest-target) {
			numClosest = numCur
		}
		// 寻路剪枝
		if numCur < target {
			// 向target移动
			root = root.Right
		} else if target < numCur {
			// 向target移动
			root = root.Left
		} else {
			return root.Val
		}
	}
	return int(numClosest)
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}
