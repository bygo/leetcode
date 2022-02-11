package main

// https://leetcode-cn.com/problems/clone-binary-tree-with-random-pointer

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

// 克隆含随机指针的二叉树

func copyRandomBinaryTree(root *Node) *NodeCopy {
	nodeMpCopy := map[*Node]*NodeCopy{}

	var dfs func(node *Node) *NodeCopy
	dfs = func(node *Node) *NodeCopy {
		if node == nil {
			return nil
		}
		_, ok := nodeMpCopy[node]
		if !ok {
			nodeMpCopy[node] = &NodeCopy{Val: node.Val}
			nodeMpCopy[node].Left = dfs(node.Left)
			nodeMpCopy[node].Right = dfs(node.Right)
			nodeMpCopy[node].Random = dfs(node.Random)
		}

		return nodeMpCopy[node]
	}
	return dfs(root)
}
