package main

// https://leetcode.cn/problems/clone-graph

type Node struct {
	Val       int
	Neighbors []*Node
}

// dfs
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	m := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		v, ok := m[n]
		if ok {
			return v
		}
		cn := &Node{Val: n.Val}
		m[n] = cn
		for i := range n.Neighbors {
			cn.Neighbors = append(cn.Neighbors, dfs(n.Neighbors[i]))
		}
		return cn
	}
	return dfs(node)
}

// bfs
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var queue = []*Node{node}
	m := map[*Node]*Node{}
	m[node] = &Node{Val: node.Val}
	for 0 < len(queue) {
		n := queue[0]
		queue = queue[1:]
		for _, neighbor := range n.Neighbors {
			_, ok := m[neighbor]
			if !ok {
				m[neighbor] = &Node{neighbor.Val, []*Node{}}
				queue = append(queue, neighbor)
			}
			m[n].Neighbors = append(m[n].Neighbors, m[neighbor])
		}
	}
	return m[node]
}
