package bfs

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的锯齿形层序遍历 

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var depsNums [][]int
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		depsNums = append(depsNums, []int{})
		top := len(depsNums) - 1
		var idxMirror int
		var step = 1
		if top%2 == 1 {
			idxMirror = queL - 1
			step = -1
		}
		for idx := 0; idx < queL; idx++ {
			node := que[idx]
			nodeMirror := que[idxMirror]
			depsNums[top] = append(depsNums[top], nodeMirror.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			idxMirror += step
		}
		que = que[queL:]
	}

	return depsNums
}
