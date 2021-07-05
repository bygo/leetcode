package inorder

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//Title: Convert Binary Search Tree to Sorted Doubly Linked List
// Link: https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list

var pre *Node
var head *Node

func treeToDoublyList(root *Node) *Node {
	pre = nil
	head = nil
	dfs(root)
	if pre != nil {
		pre.Right = head
		head.Left = pre
	}
	return head
}

func dfs(root *Node) {
	if root != nil {
		dfs(root.Left)
		if pre == nil {
			head = root
		} else {
			pre.Right = root
			root.Left = pre
		}
		pre = root
		dfs(root.Right)
	}
}
