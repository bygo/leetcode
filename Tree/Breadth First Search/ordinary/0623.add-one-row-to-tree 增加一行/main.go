package main

// https://leetcode.cn/problems/add-one-row-to-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 在二叉树中增加一行
// ⚠️ 原来层 移到 下一层

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 第一层 根节点
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	var que = []*TreeNode{root}
	var dep int

	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}
		dep++
		// 添加方式 = 修改子节点 所以需要提前
		if dep+1 == depth {
			for _, node := range que {
				node.Left = &TreeNode{Val: val, Left: node.Left}
				node.Right = &TreeNode{Val: val, Right: node.Right}
			}
			break
		}

		for _, node := range que[:cnt] {
			if node.Left != nil {
				que = append(que, node.Left)
			}

			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		que = que[cnt:]
	}
	return root
}
