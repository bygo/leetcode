package main

// https://leetcode-cn.com/problems/count-complete-tree-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 完全二叉树的节点个数

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dep int
	node := root
	for node != nil {
		node = node.Left
		dep++
	}

	idxMin := 1<<(dep-1) - 1
	idxMax := 1<<dep - 1 - 1

	exist := func(node *TreeNode, idx int) bool {
		idxLo, idxHi := idxMin, idxMax
		depCur := dep
		for 1 < depCur {
			depCur--
			idxMid := (idxLo + idxHi) >> 1
			if idxMid < idx {
				idxLo = idxMid + 1
				node = node.Right
			} else if idx <= idxMid {
				idxHi = idxMid - 1
				node = node.Left
			}
		}
		return node != nil
	}

	lo, hi := idxMin, idxMax+1
	for lo < hi {
		mid := (lo + hi) >> 1
		if exist(root, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
