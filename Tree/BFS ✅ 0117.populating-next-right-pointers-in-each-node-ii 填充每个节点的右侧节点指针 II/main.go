package main

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// ❓ 填充每个节点的下一个右侧节点指针 II

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var que = []*Node{root}
	var head = root
	var pre *Node
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		pre = nil
		for _, q := range que[:queL] {
			if pre != nil {
				pre.Next = q
			}
			pre = q
			if q.Left != nil {
				que = append(que, q.Left)
			}

			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[queL:]
	}
	return head
}
