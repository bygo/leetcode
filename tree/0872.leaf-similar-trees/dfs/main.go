package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1, l2 := []int{}, []int{}
	dfs(root1, l1)
	dfs(root2, l2)
	if len(l1) != len(l2) {
		return false
	}
	for i, _ := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, leafs []int) {
	if root != nil {
		dfs(root.Left, leafs)
		if root.Left == nil && root.Right == nil {
			leafs = append(leafs, root.Val)
		}
		dfs(root.Right, leafs)
	}
}
