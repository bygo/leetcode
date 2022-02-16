package inorder

// https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// ❓ 将二叉搜索树转化为排序的双向链表

func treeToDoublyList(root *Node) *Node {
	var pre, head *Node
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre == nil {
			head = node
		} else {
			pre.Right = node
			node.Left = pre
		}
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	if pre != nil {
		pre.Right = head
		head.Left = pre
	}
	return head
}
