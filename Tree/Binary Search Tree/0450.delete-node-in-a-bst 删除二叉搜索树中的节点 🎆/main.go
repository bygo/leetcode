package main

// https://leetcode-cn.com/problems/delete-node-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 删除二叉搜索树中的节点

func deleteNode(root *TreeNode, val int) *TreeNode {
	var del func(node *TreeNode) *TreeNode
	del = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val < val {
			// 只处理右子树
			node.Right = del(node.Right)
		} else if val < node.Val {
			// 只处理左子树
			node.Left = del(node.Left)
		} else {
			if node.Left == nil {
				// 移除节点
				node = node.Right
			} else if node.Right == nil {
				// 移除节点
				node = node.Left
			} else {

				// 最小值 替代 根节点
				//min := node.Right
				//for min.Left != nil {
				//	min = min.Left
				//}
				//min.Left = node.Left
				//node = node.Right

				// 最大值 替代 根节点
				//max := node.Left
				//for max.Right != nil {
				//	max = max.Right
				//}
				//max.Right = node.Right
				//node = node.Left

				// 替换
				max := node.Left
				for max.Right != nil {
					max = max.Right
				}
				node.Val = max.Val
				val = node.Val
				node.Left = del(node.Left)
			}
		}
		return node
	}
	return del(root)
}
