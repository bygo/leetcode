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
	var dep = -1
	node := root
	// 统计层
	for node != nil {
		node = node.Left
		dep++
	}

	// idx: -1
	// max: -1
	idxMin := 1<<dep - 1
	idxMax := 1<<(dep+1) - 1 - 1

	exist := func(node *TreeNode, idx int) bool {
		idxLo, idxHi := idxMin, idxMax
		depCur := dep
		for 0 < depCur {
			depCur--
			idxMid := (idxLo + idxHi) >> 1
			if idxMid < idx {
				idxLo = idxMid + 1
				node = node.Right
			} else if idx <= idxMid {
				// 二分时，可能刚好命中索引，必须走到底
				// 翻倍时，走路径为 left
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
