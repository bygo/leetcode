package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func inorderSuccessor(node *Node) *Node {
	if node.Right != nil { //如果有root.Right，找最小值
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	} else {
		for node.Parent != nil {
			if node.Parent.Left == node { //如果root.Right为空，寻找以本树为左子树的父节点，找到返回，找不到返回nil
				return node.Parent
			}
			node = node.Parent
		}
	}
	return nil
}
