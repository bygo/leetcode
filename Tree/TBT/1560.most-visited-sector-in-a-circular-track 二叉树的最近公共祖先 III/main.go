package main

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iii/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// ❓ 二叉树的最近公共祖先 III

func lowestCommonAncestor(p *Node, q *Node) *Node {
	var nodeMp = map[*Node]struct{}{}
	for p != nil {
		nodeMp[p] = struct{}{}
		p = p.Parent
	}

	for q != nil {
		_, ok := nodeMp[q]
		if ok {
			return q
		}
		q = q.Parent
	}
	return nil
}
