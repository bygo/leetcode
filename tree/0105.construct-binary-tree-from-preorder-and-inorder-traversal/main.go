package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val: preorder[0],
				//Val: inorder[root],
				Left:  buildTree(preorder[1:k+1], inorder[:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}
