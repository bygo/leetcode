package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
dfs
1.左节点小于根节点
2.右节点大于根节点
3.左节点的右节点小于根节点
（因为 根节点的左节点下的所有值，必须全部小于根节点）
4.右节点的左节点大于根节点
（因为 根节点的右节点下的所有值，必须全部大于根节点）
 */
func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min || max != nil && *max <= root.Val {
		return false
	}
	return dfs(root.Left, min, &root.Val) && dfs(root.Right, &root.Val, max)
}

//中序遍历
var last *TreeNode

func isValidBST2(root *TreeNode) bool {
	last = nil
	return inorder(root)
}

func inorder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !inorder(root.Left) || last != nil && root.Val <= last.Val {
		return false
	}
	last = root
	return inorder(root.Right)
}
