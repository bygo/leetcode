package main

// https://leetcode.cn/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树
// 📚 BST 节点边界

type Pair struct {
	node *TreeNode
	min  int
	max  int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var que []*Pair
	que = append(que, &Pair{
		node: root,
		min:  -1 << 63,
		max:  1<<63 - 1,
	})

	for {
		var queL = len(que)
		if queL == 0 {
			break
		}
		for _, pair := range que[:queL] {
			node := pair.node
			min := pair.min
			max := pair.max
			if node.Val <= min || max <= node.Val {
				return false
			}
			if node.Left != nil {
				que = append(que, &Pair{
					node: node.Left,
					min:  min,
					max:  node.Val,
				})
			}
			if node.Right != nil {
				que = append(que, &Pair{
					node: node.Right,
					min:  node.Val,
					max:  max,
				})
			}
		}
		que = que[queL:]
	}
	return true
}
