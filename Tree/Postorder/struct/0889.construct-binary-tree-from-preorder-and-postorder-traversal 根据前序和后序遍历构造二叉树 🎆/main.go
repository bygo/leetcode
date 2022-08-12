package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

// ❓ 根据前序和后序遍历构造二叉树

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}
	postL := len(post)
	for idx := range post {
		if post[idx] == pre[1] { // 找到左子树
			return &TreeNode{
				Val:   pre[0],
				Left:  constructFromPrePost(pre[1:idx+2], post[:idx+1]),
				Right: constructFromPrePost(pre[idx+2:], post[idx+1:postL-1]),
			}
		}
	}
	return nil
}
