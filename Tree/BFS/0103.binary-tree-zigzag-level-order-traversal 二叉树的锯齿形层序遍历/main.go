package bfs

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的锯齿形层序遍历 
// 📚 镜像遍历

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var depsNums [][]int
	var que = []*TreeNode{root}
	var dep = -1
	var idxMir, step int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		depsNums = append(depsNums, []int{})
		dep++

		if dep%2 == 1 {
			idxMir = queL - 1
			step = -1
		} else {
			idxMir = 0
			step = 1
		}

		for idx := 0; idx < queL; idx++ {
			node := que[idx]
			nodeMir := que[idxMir]
			depsNums[dep] = append(depsNums[dep], nodeMir.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			idxMir += step
		}
		que = que[queL:]
	}

	return depsNums
}
