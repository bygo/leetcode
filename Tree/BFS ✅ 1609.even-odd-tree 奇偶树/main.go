package main

// https://leetcode-cn.com/problems/even-odd-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 奇偶树
// 📚 镜像遍历

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var que = []*TreeNode{root}
	var dep, numPre int
	var idxMir, step int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		numPre = 0
		// 奇数递增 偶数递减
		var remDep = dep % 2

		// 偶数 递减
		if remDep == 1 {
			idxMir, step = queL-1, -1
		} else {
			idxMir, step = 0, 1
		}

		for idx := 0; idx < queL; idx++ {
			node := que[idx]
			nodeMir := que[idxMir]
			// 严格递增 & 索引奇偶和值奇偶 互斥
			if nodeMir.Val <= numPre || nodeMir.Val%2 == remDep {
				return false
			}
			numPre = nodeMir.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			idxMir += step
		}
		dep++
		que = que[queL:]
	}
	return true
}
