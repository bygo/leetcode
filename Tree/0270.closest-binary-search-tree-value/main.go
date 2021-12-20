package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	var res = 1<<63 - 1
	for root != nil {
		if abs(float64(root.Val)-target) < abs(float64(res)-target) {
			res = root.Val
		}
		if float64(root.Val) < target {
			root = root.Right
		} else if target < float64(root.Val) {
			root = root.Left
		} else {
			return root.Val
		}
	}
	return res
}

func abs(n float64) float64 {
	if n < 0 {
		return -1 * n
	}
	return n
}
