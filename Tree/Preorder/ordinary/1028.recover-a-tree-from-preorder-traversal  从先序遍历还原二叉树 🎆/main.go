package main

// https://leetcode.cn/problems/recover-a-tree-from-preorder-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 从先序遍历还原二叉树

func recoverFromPreorder(traversal string) *TreeNode {
	var stack []*TreeNode
	var idx int
	tL := len(traversal)
	for idx < tL {
		// 子节点层数
		cnt := 0
		for traversal[idx] == '-' {
			cnt++
			idx++
		}
		// 值
		val := 0
		for idx < tL && '0' <= traversal[idx] && traversal[idx] <= '9' {
			val = val*10 + int(traversal[idx]-'0')
			idx++
		}
		node := &TreeNode{Val: val}

		// 回到目标层数
		sL := len(stack)
		if cnt == sL {
			if 0 < sL {
				stack[sL-1].Left = node
			}
		} else {
			stack = stack[:cnt]
			stack[cnt-1].Right = node
		}
		stack = append(stack, node)
	}
	return stack[0]
}
