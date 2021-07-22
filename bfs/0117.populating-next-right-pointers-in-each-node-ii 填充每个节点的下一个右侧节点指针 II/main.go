package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// populating next right pointers in each node ii
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue = []*Node{root}
	var head = root
	var last *Node
	for 0 < len(queue) {
		var cnt = len(queue)
		last = nil
		for _, v := range queue[:cnt] {
			if last != nil {
				last.Next = v
			}
			last = v
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}
	return head
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent *Node
	parent = root

	for parent != nil {
		var pre, down *Node
		down = nil
		for parent != nil {
			if parent.Left != nil {
				if pre == nil {
					down = parent.Left
				} else {
					pre.Next = parent.Left
				}
				pre = parent.Left
			}

			if parent.Right != nil {
				if pre == nil {
					down = parent.Right
				} else {
					pre.Next = parent.Right
				}
				pre = parent.Right
			}
			parent = parent.Next
		}
		parent = down
	}
	return root
}
