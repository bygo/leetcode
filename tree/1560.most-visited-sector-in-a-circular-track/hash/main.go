package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

//Title: Most Visited Sector in  a Circular Track
//Link: https://leetcode-cn.com/problems/most-visited-sector-in-a-circular-track

func lowestCommonAncestor(p *Node, q *Node) *Node {
	var m = map[*Node]struct{}{}
	for p != nil {
		m[p] = struct{}{}
		p = p.Parent
	}

	for q != nil {
		if _, ok := m[q]; ok {
			return q
		}
		q = q.Parent
	}
	return nil
}
