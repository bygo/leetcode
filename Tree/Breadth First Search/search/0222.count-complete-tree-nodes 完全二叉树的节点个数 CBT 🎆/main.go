package main

// https://leetcode.cn/problems/count-complete-tree-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 完全二叉树的节点个数
// 📚 分区寻路

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

	check := func(node *TreeNode, idx int) bool {
		idxLo, idxHi := idxMin, idxMax
		var depCur = dep
		for 0 < depCur {
			depCur--
			//             0
			// 	     1            2
			//    3     4     5       6
			//  7  8    9 10  11 12  13  14

			idxMid := (idxLo + idxHi) >> 1
			if idxMid < idx {
				idxLo = idxMid + 1
				node = node.Right
			} else if idx < idxMid {
				idxHi = idxMid - 1
				node = node.Left
			} else if idx == idxMid {
				// 寻路时 可能刚好命中 idx
				// 命中时 因为向下取整，idx 必定 为 left 区间最大值
				// hi不再变小，lo开始逼近
				node = node.Left
				for 0 < depCur {
					depCur--
					node = node.Right
				}
				break
			}
		}
		return node != nil
	}

	lo, hi := idxMin, idxMax+1
	for lo < hi {
		mid := (lo + hi) >> 1
		if check(root, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
