package main

// https://leetcode.cn/problems/construct-binary-search-tree-from-preorder-traversal/

// ❓ 前序遍历构造二叉搜索树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
}

func bstFromPreorder(preorder []int) *TreeNode {
	preL := len(preorder)
	if preL == 0 {
		return nil
	}
	if preL == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	lo, hi := 1, preL-1
	var idx = 1
	for lo <= hi {
		idx = (lo + hi) >> 1
		if preorder[idx] < preorder[0] { // 左子树
			lo = idx + 1
		} else if preorder[0] < preorder[idx] { // 右子树
			hi = idx - 1
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:idx]),
		Right: bstFromPreorder(preorder[idx:]),
	}
}
