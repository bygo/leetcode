package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue = []*Node{root}
	head := root
	var last *Node
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		last = nil
		for _, v := range queue[:counter] {
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
		queue = queue[counter:]
	}
	return head
}
