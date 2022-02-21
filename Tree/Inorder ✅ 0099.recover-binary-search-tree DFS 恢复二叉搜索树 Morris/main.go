package main

// https://leetcode-cn.com/problems/recover-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 恢复二叉搜索树

func recoverTree(root *TreeNode) {
	var pre, first, second, max *TreeNode
	for root != nil {
		if root.Left == nil {
			if pre != nil && root.Val <= pre.Val {
				if first == nil {
					first = pre
				}
				second = root
			}
			pre = root
			root = root.Right
		} else {
			//寻找左树最大节点
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right != root {
				// 未连接,连接到root
				max.Right = root
				root = root.Left
			} else {
				// 已连接，断开连接
				max.Right = nil
				if pre != nil && root.Val <= pre.Val {
					if first == nil {
						first = pre
					}
					second = root
				}
				pre = root
				root = root.Right
			}
		}
	}
	first.Val, second.Val = second.Val, first.Val
}
