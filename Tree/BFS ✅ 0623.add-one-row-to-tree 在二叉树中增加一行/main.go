package main

// https://leetcode-cn.com/problems/add-one-row-to-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 在二叉树中增加一行

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 第一层 根节点
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	var que = []*TreeNode{root}
	var dep = 1

	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}
		dep++
		if dep == depth {
			for _, q := range que {
				q.Left = &TreeNode{Val: val, Left: q.Left}
				q.Right = &TreeNode{Val: val, Right: q.Right}
			}
			break
		}

		for _, q := range que[:cnt] {
			if q.Left != nil {
				que = append(que, q.Left)
			}

			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[cnt:]
	}
	return root
}
