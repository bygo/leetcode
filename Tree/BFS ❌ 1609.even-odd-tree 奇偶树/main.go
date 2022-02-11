package main

// https://leetcode-cn.com/problems/even-odd-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 奇偶树

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var que = []*TreeNode{root}
	var dep, numPre int
	var idxMirror, step int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		numPre = 0
		var remDep = dep % 2

		// 偶数 递减
		if remDep == 1 {
			idxMirror, step = queL-1, -1
		} else {
			idxMirror, step = 0, 1
		}

		for i := 0; i < queL; i++ {
			nodeCur := que[i]
			nodeMirror := que[idxMirror]
			// 严格递增 & 索引奇偶和值奇偶 互斥
			if nodeMirror.Val <= numPre || nodeMirror.Val%2 == remDep {
				return false
			}
			numPre = nodeMirror.Val
			idxMirror += step
			if nodeCur.Left != nil {
				que = append(que, nodeCur.Left)
			}
			if nodeCur.Right != nil {
				que = append(que, nodeCur.Right)
			}
		}
		dep++
		que = que[queL:]
	}
	return true
}
